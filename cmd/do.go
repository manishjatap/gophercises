package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/manishjagtap/taskmanager/mongopkg"
	"github.com/spf13/cobra"
)

var updateTask = func(task mongopkg.Task) error {
	return task.Update()
}

var convertStringToInt = func(num string) (int, error) {
	return strconv.Atoi(strings.TrimSpace(num))
}

//DoCommand : Mark task as completed
var DoCommand = &cobra.Command{
	Use:   "do",
	Short: "Mark task as completed",
	Run: func(c *cobra.Command, args []string) {

		taskNumber, err := convertStringToInt(args[0])

		if err != nil {
			fmt.Println(err)
			return
		}

		var dummyTask mongopkg.Task

		if tasks, err := getTask(dummyTask); err == nil {
			for i, tsk := range tasks {
				if taskNumber == (i + 1) {
					if err = updateTask(tsk); err != nil {
						fmt.Println(err)
						return
					}
				}
			}
		} else {
			fmt.Println(err)
			return
		}

		fmt.Println("Task completed!")
	},
}

func init() {
	RootCommand.AddCommand(DoCommand)
}
