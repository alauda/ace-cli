package compose

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewComposeCmd creates a new compose command.
func NewComposeCmd(alauda client.APIClient) *cobra.Command {
	composeCmd := &cobra.Command{
		Use:        "compose",
		Short:      "Manage application compose",
		Long:       ``,
		Deprecated: "and will be removed in a subsequent release",
		Hidden:     true,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	composeCmd.AddCommand(
		newUpCmd(alauda),
		newLsCmd(alauda),
		newInspectCmd(alauda),
		newPsCmd(alauda),
		newStartCmd(alauda),
		newStopCmd(alauda),
		newRmCmd(alauda),
	)

	return composeCmd
}
