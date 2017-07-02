package client

import (
	"github.com/alauda/alauda/client/rest"
)

// CreateAppTemplateData defines the request body for the CreateAppTemplate API.
type CreateAppTemplateData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateAppTemplate creates a new app template.
func (client *Client) CreateAppTemplate(data *CreateAppTemplateData, filePath string) error {
	url := client.buildURL("application-templates", "")

	request, err := client.buildCreateAppTemplateRequest(data, filePath)
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

func (client *Client) buildCreateAppTemplateRequest(data *CreateAppTemplateData, configFile string) (*rest.Request, error) {
	request := rest.NewRequest(client.Token())

	request.SetFormData(map[string]string{
		"name":        data.Name,
		"description": data.Description,
	})

	request.SetFile("template", configFile)

	return request, nil
}
