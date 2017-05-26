package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// UpdateLoadBalancerData defines the request body for the UpdateLoadBalancer API.
type UpdateLoadBalancerData struct {
	Action    string         `json:"action"`
	Listeners []ListenerData `json:"listeners"`
}

// ListenerData defines one listener in the request body for the UpdateLoadBalancer API.
type ListenerData struct {
	ServiceName   string `json:"service_name"`
	Protocol      string `json:"protocol"`
	ListenerPort  int    `json:"listener_port"`
	ContainerPort int    `json:"container_port"`
}

// UpdateLoadBalancer creates or removes listeners for the service endpoints on the load balancer.
func (client *Client) UpdateLoadBalancer(name string, data *UpdateLoadBalancerData) error {
	url := client.buildURL("load_balancers", name)

	request, err := client.buildUpdateLoadBalancerRequest(data)
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

func (client *Client) buildUpdateLoadBalancerRequest(data *UpdateLoadBalancerData) (*rest.Request, error) {
	request := rest.NewRequest(client.Token())

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request.SetBody(body)

	return request, nil
}
