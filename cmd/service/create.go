package service

import (
	"errors"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type createOptions struct {
	cluster string
	space   string
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

	return createCmd
}

func doCreate(alauda client.APIClient, name string, image string, opts *createOptions, start bool) error {
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

	data := client.CreateServiceData{
		Name:            name,
		ImageName:       imageName,
		ImageTag:        imageTag,
		Cluster:         cluster,
		Space:           space,
		TargetState:     targetState,
		InstanceSize:    "CUSTOMIZED",
		ScalingMode:     "MANUAL",
		TargetInstances: 1,
		CustomInstanceSize: client.ServiceInstanceSize{
			CPU:    0.125,
			Memory: 256,
		},
	}

	err = alauda.CreateService(&data)
	if err != nil {
		return err
	}

	return nil
}
