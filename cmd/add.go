package cmd

import (
	"fmt"
	"strings"

	"github.com/ezetter/task/db"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Long:  `Add a new task to your TODO list`,
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		err := db.AddTask(task)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}
