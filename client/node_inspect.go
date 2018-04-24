package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// InspectNode retrieves details about the specified node.
func (client *Client) InspectNode(ip string, cluster string) (*Node, error) {
	url := client.buildURL("v1", "regions", "%s/nodes/%s", cluster, ip)
	request := client.buildInspectNodeRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseInspectNodeResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildInspectNodeRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseInspectNodeResult(response *rest.Response) (*Node, error) {
	result := Node{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
