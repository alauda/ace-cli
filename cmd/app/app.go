package app

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewAppCmd creates a new app command.
func NewAppCmd(alauda client.APIClient) *cobra.Command {
	appCmd := &cobra.Command{
		Use:   "app",
		Short: "Manage applications",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	appCmd.AddCommand(
		NewLsCmd(alauda),
		NewInspectCmd(alauda),
		NewYamlCmd(alauda),
		NewStartCmd(alauda),
		NewStopCmd(alauda),
		newDeleteCmd(alauda),
		newRunCmd(alauda),
	)

	return appCmd
}
