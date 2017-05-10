package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// InspectCluster retrieves details about the specified cluster.
func (client *Client) InspectCluster(name string) (*Cluster, error) {
	url := client.buildInspectClusterURL(name)
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

func (client *Client) buildInspectClusterURL(name string) string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/regions/%s/%s", server, client.Namespace(), name)
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
