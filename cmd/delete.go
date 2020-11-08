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
	"strconv"
	"github.com/spf13/cobra"
	"github.com/kchatr/exp/todo"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Aliases: []string{"del", "remove", "rm"},
	Short: "Delete a To-Do",
	Long: `Delete a To-Do from your To-Do list. This cannot be undone.`,
	Run: deleteRun,
}

func remove(s []todo.Item, p int) []todo.Item {
	s[len(s)-1], s[p] = s[p], s[len(s)-1]
    return s[:len(s)-1]
}

func deleteRun(cmd *cobra.Command, args []string) {

	var args_int = []int{}

	if len(args) == 0 {
		log.Printf("No argument. Please enter the position of the To-Do you would like to delete.")
	}

	for _, i := range args {
		j, err := strconv.Atoi(i)

		if err != nil {
			panic(err)
		}

		j--

		args_int = append(args_int, j)
	}


	items, err := todo.ReadItems(dataFile)

	if err != nil {
		log.Printf("%v", err)
	}

	for _, i := range args_int {
		if i > len(items) {
			log.Printf("No To-Do with Position %d - please enter a number between 1 and %d", (i + 1), len(items))
			return
		}
		items = remove(items, i)
	}

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
