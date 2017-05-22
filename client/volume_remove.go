package client

import (
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// RemoveVolume deletes the specified volume.
func (client *Client) RemoveVolume(id string) error {
	url := client.buildRemoveVolumeURL(id)
	request := client.buildRemoveVolumeRequest()

	response, err := request.Delete(url)
	if err != nil {
		return err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) buildRemoveVolumeURL(id string) string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/storage/%s/volumes/%s", server, client.Namespace(), id)
}

func (client *Client) buildRemoveVolumeRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
