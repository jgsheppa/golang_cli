/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github/jgsheppa/golang_cli/todo"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	doneOpt bool
	allOpt bool
	tail bool
	head bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Long: `This allows you to view the tasks you have created for yourself.`,
	Run: listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))

	if err != nil {
		log.Printf("%v", err)
	}

	sort.Sort(todo.ByPri(items))
	
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	c := color.New(color.FgHiYellow)

	switch {
	case head:
		for i := 0; i < 5; i++ {	
				c.Fprintln(w, items[i].Label()+"\t"+items[i].PrettyDone()+"\t"+items[i].PrettyPrint()+"\t"+items[i].Text+"\t")
		}
	case tail:
		for i := len(items) - 5; i < len(items); i++ {
				c.Fprintln(w, items[i].Label()+"\t"+items[i].PrettyDone()+"\t"+items[i].PrettyPrint()+"\t"+items[i].Text+"\t")
		}
	default:
		for _, item := range items {
			if allOpt || item.Done == doneOpt {
				c.Fprintln(w, item.Label()+"\t"+item.PrettyDone()+"\t"+item.PrettyPrint()+"\t"+item.Text+"\t")
			}
	} 

	}

	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' tasks")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all tasks")
	listCmd.Flags().BoolVar(&head, "head", false, "Show first five tasks")
	listCmd.Flags().BoolVar(&tail, "tail", false, "Show last five tasks")
}
