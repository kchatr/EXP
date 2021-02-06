/*
Copyright © 2020 Kaushik Chatterjee <kchatr1729@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"
	"github.com/spf13/cobra"
	"github.com/kchatr/exp/todo"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Aliases: []string{"add", "c"},
	Short: "Create a new To-Do",
	Long: `This command creates a new To-Do item for your list.`,
	Run: addRun,
}

var priority int

/*
Run when the create command is run within the CLI application.
Uses the top-down development principles of stepwise refinement and procedural abstraction to aid in both development and maintainability.
*/
func addRun(cmd *cobra.Command, args []string) {
	
	// Selection statement run to check if there are no arguments provided
	if len(args) == 0 {
		log.Printf("No argument. Please enter the To-Do you would like to create.")
	}

	items, err := todo.ReadItems(dataFile) // Reads in the current To-Dos and assigns them into an array of type Item (the custom defined struct) 
	
	// Selection statement run to check if there was an error from reading the data
	if err != nil {
		log.Printf("%v", err)
	}

	// Iterative statement that creates a new To-Do item with a desired priority using the constructor of the struct that is then appended to items
	for _, i := range args {
		item := todo.Item{Text: i}
		item.SetPriority(priority)
		items = append(items, item)
	}

	// Selection statement that saves the items to that database and if there is any error, displays it
	if err := todo.SaveItems(dataFile, items); err != nil {
		log.Printf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Set the priority of the To-Do: 1; 2; 3")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
