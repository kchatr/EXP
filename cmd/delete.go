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
	"github.com/spf13/cobra"
	"github.com/kchatr/exp/todo"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Aliases: []string{"del", "remove", "rm"},
	Short: "Delete a To-Do from your To-Do list.",
	Long: `Delete a To-Do from your To-Do list by entering the current position of the To-Do in the List as an argument.`,
	Run: deleteRun,
}

/*
A helper method that deletes an item from an array; used to achieve procedural abstraction. 
*/
func remove(s []todo.Item, p int) []todo.Item {
	fmt.Printf("%q %v\n", s[p].Text, "deleted.") // Indicate to user that item has been deleted
	s[len(s)-1], s[p] = s[p], s[len(s)-1]
    return s[:len(s)-1]
}

/*
Run when the delete command is run within the CLI application.
Uses the top-down development principles of stepwise refinement and procedural abstraction to aid in both development and maintainability.
*/
func deleteRun(cmd *cobra.Command, args []string) {

	var args_int = []int{} // A slice (Go's dynamic array implementation) to store the arguments provided (i.e. the To-Dos to be deleted)

	// Selection statement run to check if there are no arguments provided
	if len(args) == 0 {
		log.Printf("No argument. Please enter the position of the To-Do you would like to delete.")
	}

	// Iterative statement that converts the arguments (read in as strings) to integers, ensures there is no error, and appends it to the slice
	for _, i := range args {
		j, err := strconv.Atoi(i)

		if err != nil {
			panic(err)
		}

		j--

		args_int = append(args_int, j)
	}

	// Read in the items from the local To-Do database .expdos.json
	items, err := todo.ReadItems(dataFile)

	// Selection statement run to check if there was an error from reading the data
	if err != nil {
		log.Printf("%v", err)
	}

	// Iterative statement that goes through each item in the args_int array and deletes them using the helped method remove()
	for _, i := range args_int {
		if i > len(items) {
			log.Printf("No To-Do with Position %d - please enter a number between 1 and %d", (i + 1), len(items))
			return
		}
		items = remove(items, i)
	}

	// Selection statement that saves the items to that database and if there is an error, displays it
	if err := todo.SaveItems(dataFile, items); err != nil {
		log.Printf("%v", err)
	}

}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
