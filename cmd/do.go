package cmd

import (
	"fmt"
	"strconv"

	"github.com/CodeGophercises/cli-task-manager/db"
	"github.com/CodeGophercises/cli-task-manager/utils"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	RunE: func(cmd *cobra.Command, args []string) error {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Failed to parse arg: %s\n", arg)
				continue
			}
			ids = append(ids, id)
		}

		tasks, err := db.GetAllTasks()
		if err != nil {
			return err
		}

		// Iterate through all `ids` and mark those completed
		for _, id := range ids {
			// Check id is valid
			if id <= 0 || id > len(tasks) {
				fmt.Printf("Not a valid taskID %d\n", id)
				continue
			}

			task := string(tasks[id-1].Val)
			key := utils.Itob(tasks[id-1].Id)
			err := db.Delete(key)
			if err != nil {
				fmt.Printf("Failed to mark %s as completed: %s\n", task, err)
				continue
			}

			fmt.Printf("You have completed the \"%s\" task.\n", task)
		}
		return nil

	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
