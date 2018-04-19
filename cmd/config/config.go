package config

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewConfigCmd creates a new config command.
func NewConfigCmd(alauda client.APIClient) *cobra.Command {
	configCmd := &cobra.Command{
		Use:        "config",
		Short:      "Manage configurations",
		Long:       ``,
		Deprecated: "and will be removed in a subsequent release",
		Hidden:     true,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	configCmd.AddCommand(
		newCreateCmd(alauda),
		newLsCmd(alauda),
		newInspectCmd(alauda),
		newItemsCmd(alauda),
		newRmCmd(alauda),
	)

	return configCmd
}
