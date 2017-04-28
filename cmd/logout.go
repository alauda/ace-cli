package cmd

import (
	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewLogoutCmd creates a new logout command.
func NewLogoutCmd(alauda client.AlaudaClient) *cobra.Command {
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
	viper.Set(util.SettingToken, "")

	err := util.SaveConfig()
	if err != nil {
		return err
	}

	return nil
}
