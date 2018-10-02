package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// InspectApp retrieves the details of a specific application
func (client *Client) InspectApp(cluster string, namespace string, name string) (*App, error) {
	url := client.buildURL("v2", "kubernetes", "clusters/%s/applications/%s/%s", cluster, namespace, name)
	request := client.buildInspectAppRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseInspectAppResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildInspectAppRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseInspectAppResult(response *rest.Response) (*App, error) {
	result := App{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
