package service

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type updateOptions struct {
	tag        string
	cpu        float64
	memory     int
	env        []string
	cmd        string
	entrypoint string
}

// NewUpdateCmd creates a new service update command.
func NewUpdateCmd(alauda client.APIClient) *cobra.Command {
	var opts updateOptions

	updateCmd := &cobra.Command{
		Use:   "update NAME",
		Short: "Update a service",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("service update expects NAME")
			}
			return doUpdate(alauda, args[0], &opts)
		},
	}

	updateCmd.Flags().StringVarP(&opts.tag, "tag", "t", "", "New image tag")
	updateCmd.Flags().Float64VarP(&opts.cpu, "cpu", "", 0.125, "New CPU (cores) (default: 0.125)")
	updateCmd.Flags().IntVarP(&opts.memory, "memory", "", 256, "New memory (MB) (default: 256)")
	updateCmd.Flags().StringSliceVarP(&opts.env, "env", "e", []string{}, "New environment variables")
	updateCmd.Flags().StringVarP(&opts.cmd, "run-command", "r", "", "New command to run")
	updateCmd.Flags().StringVarP(&opts.entrypoint, "entrypoint", "", "", "New entrypoint for the container")

	return updateCmd
}

func doUpdate(alauda client.APIClient, name string, opts *updateOptions) error {
	fmt.Println("[alauda] Updating", name)

	util.InitializeClient(alauda)

	err := validateResourceRequirements(opts.cpu, opts.memory)
	if err != nil {
		return err
	}

	env, err := parseEnvVars(opts.env)
	if err != nil {
		return err
	}

	data := client.UpdateServiceData{
		ImageTag:     opts.tag,
		Command:      opts.cmd,
		Entrypoint:   opts.entrypoint,
		InstanceSize: "CUSTOMIZED",
		Env:          env,
		CustomInstanceSize: client.ServiceInstanceSize{
			CPU:    opts.cpu,
			Memory: opts.memory,
		},
	}

	err = alauda.UpdateService(name, &data)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
