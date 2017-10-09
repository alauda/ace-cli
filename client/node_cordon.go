package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// CordonNodeData defines the request body for the Cordon API.
type CordonNodeData struct {
	Action string `json:"action"`
}

// CordonNode makes the specific node unschedulable.
func (client *Client) CordonNode(ip string, cluster string, data *CordonNodeData) error {
	url := client.buildURL("regions", "%s/nodes/%s/actions", cluster, ip)
	request, err := client.buildCordonNodeRequest(data)

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

func (client *Client) buildCordonNodeRequest(data *CordonNodeData) (*rest.Request, error) {
	request := rest.NewRequest(client.Token())

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request.SetBody(body)

	return request, nil
}
