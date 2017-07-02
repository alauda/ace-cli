package template

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewTemplateCmd creates a new template command.
func NewTemplateCmd(alauda client.APIClient) *cobra.Command {
	templateCmd := &cobra.Command{
		Use:   "template",
		Short: "Manage application templates",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	templateCmd.AddCommand(
		NewCreateCmd(alauda),
		NewLsCmd(alauda),
		NewInspectCmd(alauda),
		NewDownloadCmd(alauda),
	)

	return templateCmd
}
