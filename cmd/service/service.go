package service

import (
	"errors"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		NewCreateCmd(alauda),
		NewPsCmd(alauda),
		NewStartCmd(alauda),
		NewStopCmd(alauda),
		NewRmCmd(alauda),
	)

	return serviceCmd
}

func configCluster(cluster string) (string, error) {
	if cluster != "" {
		viper.Set(util.SettingCluster, cluster)

		err := util.SaveConfig()
		if err != nil {
			return "", err
		}
	}

	result := viper.GetString(util.SettingCluster)
	if result == "" {
		return "", errors.New("no cluster specified")
	}

	return result, nil
}
