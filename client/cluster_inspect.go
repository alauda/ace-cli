package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// InspectCluster retrieves details about the specified cluster.
func (client *Client) InspectCluster(name string) (*Cluster, error) {
	url := client.buildURL("v1", "regions", name)
	request := client.buildInspectClusterRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseInspectClusterResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildInspectClusterRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseInspectClusterResult(response *rest.Response) (*Cluster, error) {
	result := Cluster{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
