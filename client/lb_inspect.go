package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// InspectLoadBalancer retrieves details about a specific LB.
func (client *Client) InspectLoadBalancer(name string) (*LoadBalancer, error) {
	url := client.buildURL("v1", "load_balancers", name)
	request := client.buildInspectLoadBalancerRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseInspectLoadBalancerResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildInspectLoadBalancerRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseInspectLoadBalancerResult(response *rest.Response) (*LoadBalancer, error) {
	result := LoadBalancer{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
