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
	"fmt"
	"log"
	"strconv"
	"github.com/kchatr/exp/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Aliases: []string{"do"},
	Short: "Mark To-Do as done.",
	Long: `When you complete a To-Do on your list, run the done command to mark it as finished.`,
	Run: doneRun,
}

/*
Run when the done command is run within the CLI application.
Uses the top-down development principles of stepwise refinement and procedural abstraction to aid in both development and maintainability.
*/

func doneRun(cmd *cobra.Command, args []string) {

	// Selection statement run to check and handle the case if there are no arguments provided
	if len(args) == 0 {
		log.Println("Please enter the position of the To-Do you would like to mark as compelted.")
		return
	}
	
	items, err := todo.ReadItems(dataFile) // Returns the To-Do items and an error value from .expdos.json 
	i, err := strconv.Atoi(args[0]) // i is the position of the To-Do item to be marked done.

	// Selection statement run to check if there is an error returned by ReadItems()
	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}
	
	// Selection statement run to check if the To-Do position the user entered is valid
	if i > 0 {
		items[i-1].Done = true // Set the boolean done property of the To-Do item to true
		fmt.Printf("%q %v\n", items[i - 1].Text, "marked done.") // Indicate to user that item was marked as done.
		todo.Sort(items) // Sort the list
		todo.SaveItems(dataFile, items) // Save the items to the local database, .expdos.json
	} else {
		log.Println(i, "doesn't match any numbers")
	}

}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
