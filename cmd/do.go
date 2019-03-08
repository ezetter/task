package cmd

import (
	"fmt"
	"strconv"

	"github.com/ezetter/task/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Long:  `Mark a task on your TODO list as complete`,
	Run: func(cmd *cobra.Command, args []string) {
		taskNo, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Expected a number, not %v.\n", args[0])
			return
		}
		if len(db.ListTasks()) < taskNo || taskNo < 1 {
			fmt.Printf("There is no task #%v.\n", args[0])
			return
		}
		task := db.RemoveTask(taskNo)

		fmt.Printf("You have completed the \"%s\" task.\n", task)
	},
}
