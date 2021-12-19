/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github/jgsheppa/golang_cli/todo"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	doneOpt bool
	allOpt bool
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
	for _, item := range items {
		if allOpt || item.Done == doneOpt {
			fmt.Fprintln(w, item.Label()+"\t"+item.PrettyDone()+"\t"+item.PrettyPrint()+"\t"+item.Text+"\t")
		}

	}



	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' Todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all Todos")


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
