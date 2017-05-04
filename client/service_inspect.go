package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// InspectService starts the specified service.
func (client *Client) InspectService(name string) (*Service, error) {
	url := client.buildInspectServiceURL(name)
	request := client.buildInspectServiceRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseInspectServicesResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildInspectServiceURL(name string) string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/%s/%s/%s", server, "services", client.Namespace(), name)
}

func (client *Client) buildInspectServiceRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseInspectServicesResult(response *rest.Response) (*Service, error) {
	result := Service{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
