package cmd

import (
	"fmt"
	"strings"

	"github.com/CodeGophercises/cli-task-manager/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	RunE: func(cmd *cobra.Command, args []string) error {
		task := strings.TrimSpace(strings.Join(args, " "))
		_, err := db.CreateTask(task)
		if err != nil {
			return err
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
