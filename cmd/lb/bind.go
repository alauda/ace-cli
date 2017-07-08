package lb

import (
	"errors"
	"fmt"

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

	data, err := parseBindListeners(alauda, opts.listeners)
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

func parseBindListeners(alauda client.APIClient, listenersDesc []string) (*client.UpdateLoadBalancerData, error) {
	var listeners = []client.ListenerData{}

	for _, desc := range listenersDesc {
		serviceName, listenerPort, containerPort, protocol, err := util.ParseListener(desc)
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
		Action:    "bind",
		Listeners: listeners,
	}

	return &data, nil
}
