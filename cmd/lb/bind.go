package lb

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type bindOptions struct {
	listeners []string
}

// NewBindCmd creates a new lb bind command.
func NewBindCmd(alauda client.APIClient) *cobra.Command {
	var opts bindOptions

	bindCmd := &cobra.Command{
		Use:   "bind NAME",
		Short: "Bind service endpoints to the load balancer",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("lb bind expects NAME")
			}
			return doBind(alauda, args[0], &opts)
		},
	}

	bindCmd.Flags().StringSliceVarP(&opts.listeners, "listener", "l", []string{}, "Listeners to bind to the load balancer (serviceName:[listenerPort:]containerPort[/protocol]")

	return bindCmd
}

func doBind(alauda client.APIClient, name string, opts *bindOptions) error {
	fmt.Println("[alauda] Binding listeners to", name)

	util.InitializeClient(alauda)

	data, err := parseBindListeners(opts.listeners)
	if err != nil {
		return err
	}

	err = alauda.UpdateLoadBalancer(name, data)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}

func parseBindListeners(listenersDesc []string) (*client.UpdateLoadBalancerData, error) {
	var listeners = []client.ListenerData{}

	for _, desc := range listenersDesc {
		serviceName, listenerPort, containerPort, protocol, err := parseListener(desc)
		if err != nil {
			return nil, err
		}

		if serviceName == "" {
			return nil, errors.New("no service name specified for listener")
		}

		if protocol == "" {
			if containerPort == 80 {
				protocol = "http"
			} else {
				protocol = "tcp"
			}
		}

		listener := client.ListenerData{
			ServiceName:   serviceName,
			ListenerPort:  listenerPort,
			ContainerPort: containerPort,
			Protocol:      protocol,
		}
		listeners = append(listeners, listener)
	}

	data := client.UpdateLoadBalancerData{
		Action:    "bind",
		Listeners: listeners,
	}

	return &data, nil
}

func parseListener(desc string) (string, int, int, string, error) {
	var name string
	var listenerPort int
	var containerPort int
	var protocol string
	var err error

	result := strings.Split(desc, "/")

	if len(result) > 2 {
		return "", 0, 0, "", errors.New("invalid listener descriptor, expected [name:][listenerPort:]containerPort[/protocol]")
	}

	if len(result) == 2 {
		desc = result[0]
		protocol = result[1]

		if protocol != "http" && protocol != "tcp" {
			return "", 0, 0, "", errors.New("invalid protocol specified, supported protocols are [tcp, http]")
		}
	}

	result = strings.Split(desc, ":")

	if len(result) > 3 {
		return "", 0, 0, "", errors.New("invalid listener descriptor, expected [name:][listenerPort:]containerPort")
	}

	switch len(result) {
	case 1:
		// containerPort
		containerPort, err = strconv.Atoi(result[0])
		if err != nil {
			return "", 0, 0, "", errors.New("invalid listener descriptor, containerPort should be int")
		}
	case 2:
		// name:containerPort or listenerPort:containerPort
		containerPort, err = strconv.Atoi(result[1])
		if err != nil {
			return "", 0, 0, "", errors.New("invalid listener descriptor, expected name:containerPort or listenerPort:containerPort")
		}

		listenerPort, err = strconv.Atoi(result[0])
		if err != nil {
			name = result[0]
		}
	case 3:
		// name:listenerPort:containerPort
		name = result[0]

		listenerPort, err = strconv.Atoi(result[1])
		if err != nil {
			return "", 0, 0, "", errors.New("invalid listener descriptor, listenerPort is not int, in name:listenerPort:containerPort")
		}

		containerPort, err = strconv.Atoi(result[2])
		if err != nil {
			return "", 0, 0, "", errors.New("invalid listener descriptor, containerPort is not int, in name:listenerPort:containerPort")
		}
	}

	return name, listenerPort, containerPort, protocol, nil
}
