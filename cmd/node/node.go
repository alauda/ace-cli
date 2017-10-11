package node

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewNodeCmd creates a new node command.
func NewNodeCmd(alauda client.APIClient) *cobra.Command {
	nodeCmd := &cobra.Command{
		Use:   "node",
		Short: "Manage nodes",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	nodeCmd.AddCommand(
		newLsCmd(alauda),
		newInspectCmd(alauda),
		newCordonCmd(alauda),
		newUncordonCmd(alauda),
	)

	return nodeCmd
}
