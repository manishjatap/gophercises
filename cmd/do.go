package cmd

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/manishjagtap/taskmanager/mongopkg"
	"github.com/spf13/cobra"
)

//DoCommand : Mark task as completed
var DoCommand = &cobra.Command{
	Use:   "do",
	Short: "Mark task as completed",
	Run: func(c *cobra.Command, args []string) {

		taskNumber, err := strconv.Atoi(strings.TrimSpace(args[0]))

		if err != nil {
			log.Fatal(err)
			os.Exit(2)
		}

		dummyTask := new(mongopkg.Task)

		tasks, _ := dummyTask.Get()
		for i, tsk := range tasks {
			if taskNumber == (i + 1) {
				_ = tsk.Update()
			}
		}
	},
}

func init() {
	RootCommand.AddCommand(DoCommand)
}
