package cmd

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewLogoutCmd creates a new logout command.
func NewLogoutCmd(alauda client.APIClient) *cobra.Command {
	logoutCmd := &cobra.Command{
		Use:   "logout",
		Short: "Log out of the Alauda platform",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLogout()
		},
	}

	return logoutCmd
}

func doLogout() error {
	// TODO: Implement this.
	return nil
}
