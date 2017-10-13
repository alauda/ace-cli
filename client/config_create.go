package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// CreateConfigData defines the request body for the CreateConfig API.
type CreateConfigData struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Space       string       `json:"space_name"`
	Content     []ConfigItem `json:"content"`
}

// CreateConfig creates a configuration.
func (client *Client) CreateConfig(data *CreateConfigData) error {
	url := client.buildURL("configs", "")

	request, err := client.buildCreateConfigRequest(data)
	if err != nil {
		return err
	}

	response, err := request.Post(url)
	if err != nil {
		return err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) buildCreateConfigRequest(data *CreateConfigData) (*rest.Request, error) {
	request := rest.NewRequest(client.Token())

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request.SetBody(body)

	return request, nil
}
