package service

import (
	"errors"
	"strconv"
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
		NewRestartCmd(alauda),
		NewScaleCmd(alauda),
		NewUpdateCmd(alauda),
	)

	return serviceCmd
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

func validateResourceRequirements(cpu float64, memory int) error {
	if cpu < 0.125 || cpu > 8 {
		return errors.New("supported CPU range (cores): [0.125, 8]")
	}

	if memory < 64 || memory > 32768 {
		return errors.New("supported memory range (MB): [64, 32768]")
	}

	return nil
}

func parseScale(desc string) (string, int, error) {
	result := strings.Split(desc, "=")

	if len(result) != 2 {
		return "", 0, errors.New("invalid scale descriptor, expecting NAME=NUMBER")
	}

	name := result[0]
	number, err := strconv.Atoi(result[1])
	if err != nil {
		return "", 0, errors.New("invalid scale descriptor, expecting NAME=NUMBER （e.g. web=3)")
	}

	return name, number, nil
}

func parseName(name string) (string, string, error) {
	result := strings.Split(name, ".")

	if len(result) == 1 {
		return "", result[0], nil
	} else if len(result) == 2 {
		return result[0], result[1], nil
	} else {
		return "", "", errors.New("invalid service name, expecting \"service\" or \"app.service\"")
	}
}
