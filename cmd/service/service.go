package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewServiceCmd creates a new service command.
func NewServiceCmd(alauda client.APIClient) *cobra.Command {
	serviceCmd := &cobra.Command{
		Use:        "service",
		Short:      "Manage services",
		Long:       ``,
		Deprecated: "and will be removed in a subsequent release",
		Hidden:     true,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	serviceCmd.AddCommand(
		newCreateCmd(alauda),
		newRunCmd(alauda),
		newPsCmd(alauda),
		newStartCmd(alauda),
		newStopCmd(alauda),
		newRmCmd(alauda),
		newInspectCmd(alauda),
		newRestartCmd(alauda),
		newScaleCmd(alauda),
		newUpdateCmd(alauda),
	)

	return serviceCmd
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
