package client

import (
	"github.com/alauda/alauda/client/rest"
)

// UpdateAppTemplateData defines the request body for the UpdateAppTemplate API.
type UpdateAppTemplateData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UpdateAppTemplate updates an app template.
func (client *Client) UpdateAppTemplate(name string, data *UpdateAppTemplateData, filePath string) error {
	url := client.buildURL("application-templates", name)

	request, err := client.buildUpdateAppTemplateRequest(data, filePath)
	if err != nil {
		return err
	}

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

func (client *Client) buildUpdateAppTemplateRequest(data *UpdateAppTemplateData, configFile string) (*rest.Request, error) {
	request := rest.NewRequest(client.Token())

	request.SetFormData(map[string]string{
		"name":        data.Name,
		"description": data.Description,
	})

	request.SetFile("template", configFile)

	return request, nil
}
