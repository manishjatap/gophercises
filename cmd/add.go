package cmd

import (
	"github.com/manishjagtap/taskmanager/mongopkg"
	"github.com/spf13/cobra"
)

//AddCommand : Add a new task
var AddCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(c *cobra.Command, args []string) {
		newTask := new(mongopkg.Task) // creating new task

		for i := 0; i < len(args); i++ {
			newTask.Name += args[i] + " "
		}

		_ = newTask.Insert()
	},
}

func init() {
	RootCommand.AddCommand(AddCommand)
}
