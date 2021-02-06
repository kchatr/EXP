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
	"os"
	"text/tabwriter"
	"github.com/kchatr/exp/todo"
	"github.com/spf13/cobra"
	"github.com/olekukonko/tablewriter"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Aliases: []string{"ls"},
	Short: "List your current To-Do's",
	Long: `Listing all of the current To-Do's saved.`,
	Run: listRun,
}

// Declares the variables to be used in this file
var (
	doneFlag bool
	allFlag bool
)

/*
Run when the list command is run within the CLI application.
Uses the top-down development principles of stepwise refinement and procedural abstraction to aid in both development and maintainability.
*/
func listRun(cmd *cobra.Command, args []string) {
	
	// Items are read in using ReadItems; an example of stepwise refinement and procedural abstraction.
	items, err := todo.ReadItems(dataFile)

	var data [][]string

	// Selection statement run to check if the To-Do list is empty
	if len(items) == 0 {
		log.Println("No To-Do's in Your List - use the create command to get started!")
		return
	}

	// Selection statement run to check if there was an error from reading the data
	if err != nil {
		log.Printf("%v", err)
	} 

	// Calls Sort method created in todo.go; an example of stepwise refinement
	todo.Sort(items)

	// Iterative statement that appends all of the To-Dos in the list to a String array
	// Sequential statements are run within the FOR-EACH loop
	for _, i := range items {
		var temp []string
		temp = append(temp, i.Label())
		temp = append(temp, i.PrettyDone())
		temp = append(temp, i.PrettyPrint())
		temp = append(temp, i.Text)
		data = append(data, temp)
	}

	
	/*
	Sets the parameters for the To-Do list displayed as a table to the user. 
	Controls the appearence of the GUI.
	*/
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string {"Position", "Done?", "Priority", "Task"})

	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgHiBlueColor},
		tablewriter.Colors{tablewriter.FgWhiteColor, tablewriter.Bold, tablewriter.BgHiBlueColor},
		tablewriter.Colors{tablewriter.BgHiBlueColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgHiBlueColor, tablewriter.FgWhiteColor})

	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor})

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)

	// Iterative statement that appends all To-Do items marked done based on the condition of if either the --all or --done flag is active.
	for p, i := range data {
		if allFlag || items[p].Done == doneFlag {
			table.Append(i)
		}
	}

	// Renders the table
	table.Render()

	// Flushes the writer
	w.Flush()

}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&doneFlag, "done", false, "Show 'Done' To-Do's")
	listCmd.Flags().BoolVar(&allFlag, "all", false, "Show all To-Do's")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
