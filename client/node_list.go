package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// ListNodesResult defines the response body for the ListNodes API.
type ListNodesResult struct {
	Nodes []Node
}

// ListNodes returns all nodes in a cluster.
func (client *Client) ListNodes(cluster string) (*ListNodesResult, error) {
	url := client.buildURL("v1", "regions", "%s/nodes", cluster)
	request := client.buildListNodesRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseListNodesResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildListNodesRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseListNodesResult(response *rest.Response) (*ListNodesResult, error) {
	result := ListNodesResult{}

	err := json.Unmarshal(response.Body(), &result.Nodes)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
