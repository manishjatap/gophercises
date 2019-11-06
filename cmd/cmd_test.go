package cmd

import "testing"

func TestAddCommand(t *testing.T) {
	AddCommand.Run(addCommand, []string{"Go to gym"})
}
