package image

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewImageCmd creates a new image command.
func NewImageCmd(alauda client.APIClient) *cobra.Command {
	imageCmd := &cobra.Command{
		Use:   "image",
		Short: "Manage images",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	imageCmd.AddCommand(
		NewLsCmd(alauda),
	)

	return imageCmd
}
