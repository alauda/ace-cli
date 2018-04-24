package client

import (
	"github.com/alauda/alauda/client/rest"
)

// RemoveVolume deletes the specified volume.
func (client *Client) RemoveVolume(id string) error {
	url := client.buildURL("v1", "storage", "volumes/%s", id)
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

func (client *Client) buildRemoveVolumeRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
