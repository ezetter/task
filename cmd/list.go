package cmd

import (
	"fmt"

	"github.com/ezetter/task/db"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Long: `List all of your incomplete tasks
	`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := db.ListTasks()
		if len(tasks) == 0 {
			fmt.Println("Found no tasks. Looks like you did all of them. Good job!")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, t := range tasks {
			fmt.Printf("%d. %s\n", i+1, t)
		}
	},
}
