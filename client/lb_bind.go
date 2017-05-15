package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// BindLoadBalancerData defines the request body for the BindLoadBalancer API.
type BindLoadBalancerData struct {
	Action    string             `json:"action"`
	Listeners []BindListenerData `json:"listeners"`
}

// BindListenerData defines one listener in the request body for the BindLoadBalancer API.
type BindListenerData struct {
	ServiceName   string `json:"service_name"`
	Protocol      string `json:"protocol"`
	ListenerPort  int    `json:"listener_port"`
	ContainerPort int    `json:"container_port"`
}

// BindLoadBalancer creates listeners for the service endpoints on the load balancer.
func (client *Client) BindLoadBalancer(name string, data *BindLoadBalancerData) error {
	url := client.buildBindLoadBalancerURL(name)

	request, err := client.buildBindLoadBalancerRequest(data)
	if err != nil {
		return err
	}

	response, err := request.Put(url)
	if err != nil {
		return err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) buildBindLoadBalancerURL(name string) string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/load_balancers/%s/%s", server, client.Namespace(), name)
}

func (client *Client) buildBindLoadBalancerRequest(data *BindLoadBalancerData) (*rest.Request, error) {
	request := rest.NewRequest(client.Token())

	data.Action = "bind"

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request.SetBody(body)

	return request, nil
}
