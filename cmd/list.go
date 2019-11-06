package cmd

import (
	"fmt"

	"github.com/manishjagtap/taskmanager/mongopkg"
	"github.com/spf13/cobra"
)

//ListCommand : List all the incomplete task
var ListCommand = &cobra.Command{
	Use:   "list",
	Short: "List all the incomplete task",
	Run: func(c *cobra.Command, args []string) {
		dummyTask := new(mongopkg.Task)
		tasks, _ := dummyTask.Get()
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%v. %v\n", i+1, task.Name)
		}
	},
}

func init() {
	RootCommand.AddCommand(ListCommand)
}
