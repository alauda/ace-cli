package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

type nodeActionData struct {
	Action string `json:"action"`
}

func (client *Client) performActionOnNode(ip string, cluster string, data *nodeActionData) error {
	url := client.buildURL("v1", "regions", "%s/nodes/%s/actions", cluster, ip)
	request, err := client.buildNodeActionRequest(data)

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

func (client *Client) buildNodeActionRequest(data *nodeActionData) (*rest.Request, error) {
	request := rest.NewRequest(client.Token())

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request.SetBody(body)

	return request, nil
}
