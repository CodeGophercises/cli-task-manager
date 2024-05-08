package cmd

import (
	"fmt"

	"github.com/CodeGophercises/cli-task-manager/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := db.GetAllTasks()
		if err != nil {
			return err
		}
		if len(tasks) == 0 {
			fmt.Println("You have no unfinished business. Way to go man!")
			return nil
		}
		fmt.Printf("You have the following tasks:\n")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Val)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
