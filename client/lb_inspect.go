package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// InspectLoadBalancer retrieves details about a specific LB.
func (client *Client) InspectLoadBalancer(name string) (*LoadBalancer, error) {
	url := client.buildInspectLoadBalancerURL(name)
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

func (client *Client) buildInspectLoadBalancerURL(name string) string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/load_balancers/%s/%s", server, client.Namespace(), name)
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
