package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// InspectService retrieves details about the specified service.
func (client *Client) InspectService(name string) (*Service, error) {
	url := client.buildURL("services", name)
	request := client.buildInspectServiceRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseInspectServiceResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildInspectServiceRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseInspectServiceResult(response *rest.Response) (*Service, error) {
	result := Service{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
