package service

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewServiceCmd creates a new service command.
func NewServiceCmd(alauda client.APIClient) *cobra.Command {
	serviceCmd := &cobra.Command{
		Use:   "service",
		Short: "Manage services",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	serviceCmd.AddCommand(
		NewPsCmd(alauda),
		NewStartCmd(alauda),
		NewStopCmd(alauda),
	)

	return serviceCmd
}
