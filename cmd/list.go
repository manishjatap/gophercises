package cmd

import (
	"fmt"

	"github.com/manishjagtap/taskmanager/mongopkg"
	"github.com/spf13/cobra"
)

var getTask = func(task mongopkg.Task) ([]mongopkg.Task, error) {
	return task.Get()
}

//ListCommand : List all the incomplete task
var ListCommand = &cobra.Command{
	Use:   "list",
	Short: "List all the incomplete task",
	Run: func(c *cobra.Command, args []string) {
		var dummyTask mongopkg.Task
		if tasks, err := getTask(dummyTask); err == nil {
			fmt.Println("You have the following tasks:")
			for i, task := range tasks {
				fmt.Printf("%v. %v\n", i+1, task.Name)
			}
		} else {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	RootCommand.AddCommand(ListCommand)
}
