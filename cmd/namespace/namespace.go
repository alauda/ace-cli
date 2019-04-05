package namespace

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewNamespaceCmd creates a new namespace command.
func NewNamespaceCmd(alauda client.APIClient) *cobra.Command {
	namespaceCmd := &cobra.Command{
		Use:   "namespace",
		Short: "Manage namespaces",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	namespaceCmd.AddCommand(
		newLsCmd(alauda),
		newGetCmd(alauda),
		newSetCmd(alauda),
		newInspectCmd(alauda),
	)

	return namespaceCmd
}
