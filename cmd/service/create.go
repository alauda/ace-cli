package service

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type createOptions struct {
	cluster    string
	space      string
	expose     []int
	cpu        float64
	memory     int
	number     int
	env        []string
	cmd        string
	entrypoint string
	publish    []string
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
	createCmd.Flags().StringVarP(&opts.cmd, "run-command", "r", "", "Command to run")
	createCmd.Flags().StringVarP(&opts.entrypoint, "entrypoint", "", "", "Entrypoint for the container")
	createCmd.Flags().StringSliceVarP(&opts.publish, "publish", "p", []string{}, "Ports to publish on the load balancer ([lb:][listenerPort:]containerPort[/protocol]")

	return createCmd
}

func doCreate(alauda client.APIClient, name string, image string, opts *createOptions, start bool) error {
	fmt.Println("[alauda] Creating", name)

	util.InitializeClient(alauda)

	imageName, imageTag, err := util.ParseImageNameTag(image)
	if err != nil {
		return err
	}

	opts.cluster, err = util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	opts.space, err = configSpace(opts.space)
	if err != nil {
		return err
	}

	targetState := "STOPPED"

	if start {
		targetState = "STARTED"
	}

	err = validateResourceRequirements(opts.cpu, opts.memory)
	if err != nil {
		return err
	}

	env, err := parseEnvVars(opts.env)
	if err != nil {
		return err
	}

	loadBalancers, err := parsePublish(alauda, opts)
	if err != nil {
		return err
	}

	data := client.CreateServiceData{
		Version:         "v2",
		Name:            name,
		ImageName:       imageName,
		ImageTag:        imageTag,
		Command:         opts.cmd,
		Entrypoint:      opts.entrypoint,
		Cluster:         opts.cluster,
		Space:           opts.space,
		TargetState:     targetState,
		InstanceSize:    "CUSTOMIZED",
		ScalingMode:     "MANUAL",
		TargetInstances: opts.number,
		Ports:           opts.expose,
		NetworkMode:     "BRIDGE",
		Env:             env,
		LoadBalancers:   loadBalancers,
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

func parsePublish(alauda client.APIClient, opts *createOptions) ([]client.ServiceLoadBalancer, error) {
	var loadBalancers = []client.ServiceLoadBalancer{}
	lbMap := make(map[string]*client.ServiceLoadBalancer)

	for _, desc := range opts.publish {
		lbName, listenerPort, containerPort, protocol, err := util.ParseListener(desc)
		if err != nil {
			return nil, err
		}

		var lbID string
		if lbName == "" {
			lbID, err = getDefaultLoadBalancerID(alauda, opts.cluster)
			if err != nil {
				return nil, err
			}
		} else {
			lbID, err = getLoadBalancerID(alauda, lbName)
			if err != nil {
				return nil, err
			}
		}

		if protocol == "" {
			if containerPort == 80 {
				protocol = "http"
			} else {
				protocol = "tcp"
			}
		}

		listener := client.ServiceListener{
			ListenerPort:  listenerPort,
			ContainerPort: containerPort,
			Protocol:      protocol,
		}

		if lbMap[lbName] == nil {
			lbMap[lbName] = &client.ServiceLoadBalancer{
				ID:   lbID,
				Type: "haproxy",
			}
		}

		lb := lbMap[lbName]
		lb.Listeners = append(lb.Listeners, listener)
	}

	for _, lb := range lbMap {
		loadBalancers = append(loadBalancers, *lb)
	}

	return loadBalancers, nil
}

func getLoadBalancerID(alauda client.APIClient, name string) (string, error) {
	lb, err := alauda.InspectLoadBalancer(name)
	if err != nil {
		return "", err
	}

	return lb.ID, nil
}

func getDefaultLoadBalancerID(alauda client.APIClient, cluster string) (string, error) {
	params := client.ListLoadBalancersParams{
		Cluster: cluster,
		Service: "",
	}

	result, err := alauda.ListLoadBalancers(&params)
	if err != nil {
		return "", err
	}

	for _, lb := range result.LoadBalancers {
		if lb.Type == "haproxy" && lb.AddressType == "external" {
			return lb.ID, nil
		}
	}

	return "", errors.New("no external haproxy found")
}
