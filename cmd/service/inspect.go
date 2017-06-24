package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewInspectCmd creates a new service inspect command.
func NewInspectCmd(alauda client.APIClient) *cobra.Command {
	inspectCmd := &cobra.Command{
		Use:   "inspect NAME",
		Short: "Inspect a service",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("service inspect expects NAME")
			}
			return doInspect(alauda, args[0])
		},
	}

	return inspectCmd
}

func doInspect(alauda client.APIClient, name string) error {
	fmt.Println("[alauda] Inspecting", name)

	util.InitializeClient(alauda)

	appName, serviceName, err := parseName(name)
	if err != nil {
		return err
	}

	params := client.InspectServiceParams{
		App: "",
	}

	if appName != "" {
		params.App = appName
	}

	result, err := alauda.InspectService(serviceName, &params)
	if err != nil {
		return err
	}

	err = util.Print(result)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
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
