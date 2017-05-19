package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// ListStacksParams defines the query parameters for the ListStacks API.
type ListStacksParams struct {
	Cluster string
}

// ListStacksResult defines the response body for the ListStacks API.
type ListStacksResult struct {
	Stacks []Stack
}

// ListStacks returns all stacks in a cluster.
func (client *Client) ListStacks(params *ListStacksParams) (*ListStacksResult, error) {
	url := client.buildListStacksURL()
	request := client.buildListStacksRequest(params)

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseListStacksResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildListStacksURL() string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/applications/%s", server, client.Namespace())
}

func (client *Client) buildListStacksRequest(params *ListStacksParams) *rest.Request {
	request := rest.NewRequest(client.Token())

	if params.Cluster != "" {
		request.SetQueryParam("region", params.Cluster)
	}

	return request
}

func parseListStacksResult(response *rest.Response) (*ListStacksResult, error) {
	result := ListStacksResult{}

	err := json.Unmarshal(response.Body(), &result.Stacks)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
