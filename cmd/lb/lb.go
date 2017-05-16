package lb

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewLBCmd creates a new lb command.
func NewLBCmd(alauda client.APIClient) *cobra.Command {
	lbCmd := &cobra.Command{
		Use:   "lb",
		Short: "Manage load balancers",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	lbCmd.AddCommand(
		NewLsCmd(alauda),
		NewInspectCmd(alauda),
		NewBindCmd(alauda),
		NewUnbindCmd(alauda),
	)

	return lbCmd
}
