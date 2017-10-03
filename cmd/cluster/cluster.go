package cluster

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewClusterCmd creates a new cluster command.
func NewClusterCmd(alauda client.APIClient) *cobra.Command {
	clusterCmd := &cobra.Command{
		Use:   "cluster",
		Short: "Manage clusters",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	clusterCmd.AddCommand(
		newLsCmd(alauda),
		newInspectCmd(alauda),
		newGetCmd(alauda),
		newSetCmd(alauda),
	)

	return clusterCmd
}
