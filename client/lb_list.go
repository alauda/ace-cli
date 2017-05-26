package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// ListLoadBalancersParams defines the query parameters for the ListLoadBalancers API.
type ListLoadBalancersParams struct {
	Cluster string
	Service string
}

// ListLoadBalancersResult defines the response body for the ListLoadBalancers API.
type ListLoadBalancersResult struct {
	LoadBalancers []LoadBalancer
}

// ListLoadBalancers returns all LBs in a cluster, potentially filtered by a specific service.
func (client *Client) ListLoadBalancers(params *ListLoadBalancersParams) (*ListLoadBalancersResult, error) {
	url := client.buildURL("load_balancers", "")
	request := client.buildListLoadBalancersRequest(params)

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseListLoadBalancersResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildListLoadBalancersRequest(params *ListLoadBalancersParams) *rest.Request {
	request := rest.NewRequest(client.Token())

	if params.Cluster != "" {
		request.SetQueryParam("region_name", params.Cluster)
	}

	request.SetQueryParam("detail", "true")

	if params.Service != "" {
		request.SetQueryParam("service_name", params.Service)
	}

	return request
}

func parseListLoadBalancersResult(response *rest.Response) (*ListLoadBalancersResult, error) {
	result := ListLoadBalancersResult{}

	err := json.Unmarshal(response.Body(), &result.LoadBalancers)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
