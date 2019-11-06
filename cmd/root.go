package cmd

import "github.com/spf13/cobra"

//RootCommand : Base command
var RootCommand = &cobra.Command{
	Use:   "taskmanager",
	Short: "Task handling utility",
}
