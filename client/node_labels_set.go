package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// SetNodeLabelsData defines the request body for the SetNodeLabels API.
type SetNodeLabelsData map[string]string

// SetNodeLabels sets labels on the specified node.
func (client *Client) SetNodeLabels(ip string, cluster string, data *SetNodeLabelsData) error {
	url := client.buildURL("v1", "regions", "%s/nodes/%s/labels", cluster, ip)
	request, err := client.buildSetNodeLabelsRequest(data)

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

func (client *Client) buildSetNodeLabelsRequest(data *SetNodeLabelsData) (*rest.Request, error) {
	request := rest.NewRequest(client.Token())

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request.SetBody(body)

	return request, nil
}
