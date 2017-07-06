package config

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewConfigCmd creates a new config command.
func NewConfigCmd(alauda client.APIClient) *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Manage configurations",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	configCmd.AddCommand(
		NewLsCmd(alauda),
		NewInspectCmd(alauda),
		NewItemsCmd(alauda),
	)

	return configCmd
}
