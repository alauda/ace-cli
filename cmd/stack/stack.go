package stack

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewStackCmd creates a new stack command.
func NewStackCmd(alauda client.APIClient) *cobra.Command {
	stackCmd := &cobra.Command{
		Use:   "stack",
		Short: "Manage stacks",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	stackCmd.AddCommand(
		NewLsCmd(alauda),
	)

	return stackCmd
}
