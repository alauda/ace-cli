package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// InspectVolume retrieves details about a specific volume.
func (client *Client) InspectVolume(id string) (*Volume, error) {
	url := client.buildInspectVolumeURL(id)
	request := client.buildInspectVolumeRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseInspectVolumeResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildInspectVolumeURL(id string) string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/storage/%s/volumes/%s", server, client.Namespace(), id)
}

func (client *Client) buildInspectVolumeRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseInspectVolumeResult(response *rest.Response) (*Volume, error) {
	result := Volume{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
