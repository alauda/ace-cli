package service

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type createOptions struct {
	cluster string
	space   string
	expose  []int
	cpu     float64
	memory  int
	number  int
	env     []string
}

// NewCreateCmd creates a new create service command.
func NewCreateCmd(alauda client.APIClient) *cobra.Command {
	var opts createOptions

	createCmd := &cobra.Command{
		Use:   "create NAME IMAGE",
		Short: "Create a new service",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("service create expects NAME and IMAGE")
			}
			return doCreate(alauda, args[0], args[1], &opts, false)
		},
	}

	createCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster to create the service in")
	createCmd.Flags().StringVarP(&opts.space, "space", "s", "", "Space to create the service in")
	createCmd.Flags().IntSliceVarP(&opts.expose, "expose", "", []int{}, "Ports exposed")
	createCmd.Flags().Float64VarP(&opts.cpu, "cpu", "", 0.125, "CPU (cores) (default: 0.125)")
	createCmd.Flags().IntVarP(&opts.memory, "memory", "", 256, "Memory (MB) (default: 256)")
	createCmd.Flags().IntVarP(&opts.number, "num-instances", "n", 1, "Number of instances (default: 1)")
	createCmd.Flags().StringSliceVarP(&opts.env, "env", "e", []string{}, "Environment variables")

	return createCmd
}

func doCreate(alauda client.APIClient, name string, image string, opts *createOptions, start bool) error {
	fmt.Println("[alauda] Creating", name)

	util.InitializeClient(alauda)

	imageName, imageTag, err := util.ParseImageNameTag(image)
	if err != nil {
		return err
	}

	cluster, err := configCluster(opts.cluster)
	if err != nil {
		return err
	}

	space, err := configSpace(opts.space)
	if err != nil {
		return err
	}

	targetState := "STOPPED"

	if start {
		targetState = "STARTED"
	}

	err = validateResourceRequirements(opts)
	if err != nil {
		return err
	}

	env, err := parseEnvVars(opts)
	if err != nil {
		return err
	}

	data := client.CreateServiceData{
		Version:         "v2",
		Name:            name,
		ImageName:       imageName,
		ImageTag:        imageTag,
		Cluster:         cluster,
		Space:           space,
		TargetState:     targetState,
		InstanceSize:    "CUSTOMIZED",
		ScalingMode:     "MANUAL",
		TargetInstances: opts.number,
		Ports:           opts.expose,
		NetworkMode:     "BRIDGE",
		Env:             env,
		CustomInstanceSize: client.ServiceInstanceSize{
			CPU:    opts.cpu,
			Memory: opts.memory,
		},
	}

	err = alauda.CreateService(&data)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
