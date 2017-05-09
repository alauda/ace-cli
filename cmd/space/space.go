package space

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewSpaceCmd creates a new space command.
func NewSpaceCmd(alauda client.APIClient) *cobra.Command {
	spaceCmd := &cobra.Command{
		Use:   "space",
		Short: "Manage spaces",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	spaceCmd.AddCommand(
		NewLsCmd(alauda),
	)

	return spaceCmd
}
