package registry

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewRegistryCmd creates a new registry command.
func NewRegistryCmd(alauda client.APIClient) *cobra.Command {
	registryCmd := &cobra.Command{
		Use:   "registry",
		Short: "Manage registries",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	registryCmd.AddCommand(
		newLsCmd(alauda),
		newProjectsCmd(alauda),
	)

	return registryCmd
}
