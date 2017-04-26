package cmd

import (
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of the Alauda platform",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return doLogout()
	},
}

func init() {
	RootCmd.AddCommand(logoutCmd)
}

func doLogout() error {
	// TODO: Add SDK integration here.
	return nil
}
