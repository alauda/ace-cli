package template

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type updateOptions struct {
	filePath string
}

// NewUpdateCmd creates a new template update command.
func NewUpdateCmd(alauda client.APIClient) *cobra.Command {
	var opts updateOptions

	updateCmd := &cobra.Command{
		Use:   "update NAME",
		Short: "Updates an app template",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("template update expects NAME")
			}
			return doUpdate(alauda, args[0], &opts)
		},
	}

	updateCmd.Flags().StringVarP(&opts.filePath, "file", "f", "./alauda-compose.yml", "Compose yaml file template")

	return updateCmd
}

func doUpdate(alauda client.APIClient, name string, opts *updateOptions) error {
	fmt.Printf("[alauda] Updating %s using %s\n", name, opts.filePath)

	util.InitializeClient(alauda)

	data := client.UpdateAppTemplateData{
		Name:        name,
		Description: "",
	}

	absPath, err := filepath.Abs(opts.filePath)
	if err != nil {
		return err
	}

	err = alauda.UpdateAppTemplate(name, &data, absPath)
	if err != nil {
		return err
	}

	return nil
}
