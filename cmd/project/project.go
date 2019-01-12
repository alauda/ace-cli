package project

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewProjectCmd creates a new project command.
func NewProjectCmd(alauda client.APIClient) *cobra.Command {
	projectCmd := &cobra.Command{
		Use:   "project",
		Short: "Manage projects",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	projectCmd.AddCommand(
		newGetCmd(alauda),
		newSetCmd(alauda),
		newLsCmd(alauda),
		newInspectCmd(alauda),
	)

	return projectCmd
}
