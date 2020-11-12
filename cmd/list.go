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
	// "fmt"
	"log"
	// "sort"
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

var (
	doneFlag bool
	allFlag bool
)

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)

	var data [][]string

	if len(items) == 0 {
		log.Println("No To-Do's in Your List - use the create command to get started!")
		return
	}

	if err != nil {
		log.Printf("%v", err)
	} 

	// sort.Sort(todo.ByPri(items))
	todo.Sort(items)

	for _, i := range items {
		var temp []string
		temp = append(temp, i.Label())
		temp = append(temp, i.PrettyDone())
		temp = append(temp, i.PrettyPrint())
		temp = append(temp, i.Text)
		data = append(data, temp)
	}

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

	// fmt.Fprintln(w, "Position" + "\t" + "Done?" + "\t" + "Priority" + "\t" + "Task")

	// for _, i := range items {
	// 	if allFlag || i.Done == doneFlag {
	// 		fmt.Fprintln(w, i.Label() + "\t" + i.PrettyDone() + "\t" + i.PrettyPrint() + "\t" + i.Text + "\t")
	// 	}
	// }

	for p, i := range data {
		if allFlag || items[p].Done == doneFlag {
			table.Append(i)
		}
	}

	table.Render()

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
