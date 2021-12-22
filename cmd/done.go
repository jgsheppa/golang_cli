/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github/jgsheppa/golang_cli/todo"
	"log"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark tasks as done.",
	Long: `This will allow you to mark a task as done, in order to keep better track of your todos.`,
	Run: runDone,
}

func runDone(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	i, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln(args[0], "is not a valid label.\n Make sure you enclose list items in quotes.\n", err)
	}

	if i > 0 && i < len(items) {
		items[i -1].Done = true
		fmt.Printf("%q %v\n", items[i-1].Text, "Marked Done")

		sort.Sort(todo.ByPri(items))
		todo.SaveItems(viper.GetString("datafile"), items)
	} else {
		log.Println(i, "doesn't match any items")
	}
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
