package service

import (
	"errors"
	"strings"

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
		NewRunCmd(alauda),
		NewPsCmd(alauda),
		NewStartCmd(alauda),
		NewStopCmd(alauda),
		NewRmCmd(alauda),
		NewInspectCmd(alauda),
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

func configSpace(space string) (string, error) {
	if space != "" {
		viper.Set(util.SettingSpace, space)

		err := util.SaveConfig()
		if err != nil {
			return "", err
		}
	}

	result := viper.GetString(util.SettingSpace)
	if result == "" {
		return "", errors.New("no space specified")
	}

	return result, nil
}

func validateResourceRequirements(opts *createOptions) error {
	if opts.cpu < 0.125 || opts.cpu > 8 {
		return errors.New("supported CPU range (cores): [0.125, 8]")
	}

	if opts.memory < 64 || opts.memory > 32768 {
		return errors.New("supported memory range (MB): [64, 32768]")
	}

	return nil
}

func parseEnvVars(opts *createOptions) (map[string]string, error) {
	envvars := make(map[string]string)

	for _, desc := range opts.env {
		k, v, err := parseEnvVar(desc)
		if err != nil {
			return nil, err
		}
		envvars[k] = v
	}

	return envvars, nil
}

func parseEnvVar(desc string) (string, string, error) {
	result := strings.Split(desc, "=")

	if len(result) != 2 {
		return "", "", errors.New("invalid environment variable descriptor")
	}

	return result[0], result[1], nil
}
