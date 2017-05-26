package compose

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewComposeCmd creates a new compose command.
func NewComposeCmd(alauda client.APIClient) *cobra.Command {
	composeCmd := &cobra.Command{
		Use:   "compose",
		Short: "Manage application compose",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	composeCmd.AddCommand(
		NewLsCmd(alauda),
		NewInspectCmd(alauda),
		NewPsCmd(alauda),
		NewStartCmd(alauda),
	)

	return composeCmd
}
