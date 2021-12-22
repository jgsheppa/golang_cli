/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github/jgsheppa/golang_cli/todo"
	"log"
	"sort"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task from your list",
	Long: `Deletes a task from your list by task number or exact task name`,
	Run: runDelete,
}

func runDelete (cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	i, err := strconv.Atoi(args[0])

	c := color.New(color.FgHiRed)

	if err != nil {
		log.Fatalln(args[0], "is not a valid task number.\n Make sure you enclose list items in quotes.\n", err)
	}

	if i > 0 && i < len(items) {
		itemToDelete := items[i -1]
		items[i -1] = items[len(items) - 1]
		items = items[:len(items)-1]
		c.Printf("%q %v\n", itemToDelete.Text, "Deleted")

		sort.Sort(todo.ByPri(items))
		todo.SaveItems(viper.GetString("datafile"), items)
	} else {
		log.Println(i, "doesn't match any task number")
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)

}
