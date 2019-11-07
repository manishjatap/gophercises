package cmd

import (
	"fmt"

	"github.com/manishjagtap/taskmanager/mongopkg"
	"github.com/spf13/cobra"
)

var insertTask = func(newTask mongopkg.Task) error {
	return newTask.Insert()
}

//AddCommand : Add a new task
var AddCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(c *cobra.Command, args []string) {
		var newTask mongopkg.Task // creating new task

		for i := 0; i < len(args); i++ {
			newTask.Name += args[i] + " "
		}

		err := insertTask(newTask)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Task added!")
	},
}

func init() {
	RootCommand.AddCommand(AddCommand)
}
