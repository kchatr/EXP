/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	Short: "Create a new To-Do",
	Long: `This command creates a new To-Do item for your list.`,
	Run: addRun,
}

var priority int

func addRun(cmd *cobra.Command, args []string) {

	items, err := todo.ReadItems("C:/Users/cha_k/.expdos.json") // An array of To-Do items
	
	if err != nil {
		log.Printf("%v", err)
	}

	for _, i := range args{
		item := todo.Item{Text: i}
		item.SetPriority(priority)
		items = append(items, item)
	}

	err = todo.SaveItems(dataFile, items)
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Set the priority of the To-Do:1,2,3")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
