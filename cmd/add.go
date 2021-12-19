/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github/jgsheppa/golang_cli/todo"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Long: `Add will create a new todo item for the list.`,
	Run: addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}
	for _, task := range args {
		item := todo.Item{Text: task}
		item.SetPriority(priority)
		items = append(items, item)
	}
if err := todo.SaveItems(viper.GetString("datafile"), items); err != nil {
	fmt.Errorf("%v", err)
}

}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
