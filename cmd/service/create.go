package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/alauda/alauda/cmd/volume"
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
	volumes    []string
	configs    []string
}

// NewCreateCmd creates a new service create command.
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
	createCmd.Flags().StringSliceVarP(&opts.volumes, "volume", "v", []string{}, "Volumes to mount to the container (volumeName:containerPath or hostPath:containerPath)")
	createCmd.Flags().StringSliceVarP(&opts.configs, "config", "", []string{}, "Configuration to inject into the container (name:key:path)")

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

	env, err := util.ParseKeyValues(opts.env)
	if err != nil {
		return err
	}

	loadBalancers, published, err := parsePublish(alauda, opts)
	if err != nil {
		return err
	}

	ports := append(opts.expose, published...)

	volumes, err := parseVolumes(alauda, opts)
	if err != nil {
		return err
	}

	configs, err := parseConfigs(alauda, opts)
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
		Ports:           ports,
		NetworkMode:     "BRIDGE",
		Env:             env,
		LoadBalancers:   loadBalancers,
		Volumes:         volumes,
		Configs:         configs,
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

func parsePublish(alauda client.APIClient, opts *createOptions) ([]client.ServiceLoadBalancer, []int, error) {
	var ports = []int{}
	var loadBalancers = []client.ServiceLoadBalancer{}
	lbMap := make(map[string]*client.ServiceLoadBalancer)

	for _, desc := range opts.publish {
		lbName, listenerPort, containerPort, protocol, err := util.ParseListener(desc)
		if err != nil {
			return nil, nil, err
		}

		ports = append(ports, containerPort)

		var lbID string
		if lbName == "" {
			lbID, err = getDefaultLoadBalancerID(alauda, opts.cluster)
			if err != nil {
				return nil, nil, err
			}
		} else {
			lbID, err = getLoadBalancerID(alauda, lbName)
			if err != nil {
				return nil, nil, err
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

		if lbMap[lbID] == nil {
			lbMap[lbID] = &client.ServiceLoadBalancer{
				ID:   lbID,
				Type: "haproxy",
			}
		}

		lb := lbMap[lbID]
		lb.Listeners = append(lb.Listeners, listener)
	}

	for _, lb := range lbMap {
		loadBalancers = append(loadBalancers, *lb)
	}

	return loadBalancers, ports, nil
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

func parseVolumes(alauda client.APIClient, opts *createOptions) ([]client.ServiceVolume, error) {
	var volumes = []client.ServiceVolume{}

	for _, desc := range opts.volumes {
		volumeName, containerPath, err := parseVolume(desc)
		if err != nil {
			return nil, err
		}

		volumeID := "host_path"

		if !strings.HasPrefix(volumeName, "/") {
			volumeID, err = volume.GetVolumeID(alauda, volumeName)
			if err != nil {
				return nil, err
			}
		}

		volume := client.ServiceVolume{
			Path:       containerPath,
			VolumeName: volumeName,
			VolumeID:   volumeID,
		}

		volumes = append(volumes, volume)
	}

	return volumes, nil
}

func parseVolume(desc string) (string, string, error) {
	result := strings.Split(desc, ":")

	if len(result) != 2 {
		return "", "", errors.New("invalid volume descriptor, expecting volumeName:containerPath or hostPath:containerPath")
	}

	return result[0], result[1], nil
}

func parseConfigs(alauda client.APIClient, opts *createOptions) ([]client.ServiceConfig, error) {
	var configs = []client.ServiceConfig{}

	for _, desc := range opts.configs {
		name, key, path, err := parseConfig(desc)
		if err != nil {
			return nil, err
		}

		config := client.ServiceConfig{
			Type: "config",
			Path: path,
			Value: client.ServiceConfigValue{
				Name: name,
				Key:  key,
			},
		}

		configs = append(configs, config)
	}

	return configs, nil
}

func parseConfig(desc string) (string, string, string, error) {
	result := strings.Split(desc, ":")

	if len(result) != 3 {
		return "", "", "", errors.New("invalid config descriptor, expecting name:key:path")
	}

	return result[0], result[1], result[2], nil
}
