package lb

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

func newUnbindCmd(alauda client.APIClient) *cobra.Command {
	var opts bindOptions

	unbindCmd := &cobra.Command{
		Use:   "unbind NAME",
		Short: "Remove the bindings between service endpoints and the load balancer",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("lb unbind expects NAME")
			}
			return doUnbind(alauda, args[0], &opts)
		},
	}

	unbindCmd.Flags().StringSliceVarP(&opts.listeners, "listener", "l", []string{}, "Listeners to remove from the load balancer (servicename:listenerPort:containerPort")

	return unbindCmd
}

func doUnbind(alauda client.APIClient, name string, opts *bindOptions) error {
	fmt.Println("[alauda] Unbinding listeners from", name)

	util.InitializeClient(alauda)

	data, err := parseUnbindListeners(alauda, opts.listeners)
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

func parseUnbindListeners(alauda client.APIClient, listenersDesc []string) (*client.UpdateLoadBalancerData, error) {
	var listeners = []client.ListenerData{}

	for _, desc := range listenersDesc {
		serviceName, listenerPort, containerPort, protocol, err := util.ParseListener(desc)
		if err != nil {
			return nil, err
		}

		if serviceName == "" {
			return nil, errors.New("no service name specified for listener")
		}

		if listenerPort == 0 {
			return nil, errors.New("no listener port specified for listener")
		}

		if protocol != "" {
			return nil, errors.New("invalid listener descriptor, expected serviceName:listenerPort:containerPort")
		}

		params := client.ServiceParams{
			App: "",
		}

		service, err := alauda.InspectService(serviceName, &params)
		if err != nil {
			return nil, err
		}

		listener := client.ListenerData{
			ServiceID:     service.ID,
			ListenerPort:  listenerPort,
			ContainerPort: containerPort,
			Protocol:      protocol,
		}
		listeners = append(listeners, listener)
	}

	data := client.UpdateLoadBalancerData{
		Action:    "unbind",
		Listeners: listeners,
	}

	return &data, nil
}
