package service

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

func newScaleCmd(alauda client.APIClient) *cobra.Command {
	scaleCmd := &cobra.Command{
		Use:   "scale NAME=NUMBER",
		Short: "Scale a service to the specified number of instances",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("service start expects NAME=NUMBER")
			}
			return doScale(alauda, args[0])
		},
	}

	return scaleCmd
}

func doScale(alauda client.APIClient, desc string) error {
	name, number, err := parseScale(desc)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] Scaling", name)

	util.InitializeClient(alauda)

	appName, serviceName, err := parseName(name)
	if err != nil {
		return err
	}

	params := client.ServiceParams{
		App: "",
	}

	if appName != "" {
		params.App = appName
	}

	data := client.ScaleServiceData{
		TargetInstances: number,
	}

	err = alauda.ScaleService(serviceName, &data, &params)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
