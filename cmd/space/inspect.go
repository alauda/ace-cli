package space

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type inspectOptions struct {
	project string
}

func newInspectCmd(alauda client.APIClient) *cobra.Command {
	var opts inspectOptions

	inspectCmd := &cobra.Command{
		Use:   "inspect NAME",
		Short: "Inspect a space",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("space inspect expects NAME")
			}
			return doInspect(alauda, args[0], &opts)
		},
	}

	inspectCmd.Flags().StringVarP(&opts.project, "project", "p", "", "Project")

	return inspectCmd
}

func doInspect(alauda client.APIClient, name string, opts *inspectOptions) error {
	fmt.Println("[alauda] Inspecting", name)

	util.InitializeClient(alauda)

	project, err := util.ConfigProject(opts.project)
	if err != nil {
		return err
	}

	params := client.InspectSpaceParams{
		Project: project,
	}

	result, err := alauda.InspectSpace(name, &params)
	if err != nil {
		return err
	}

	err = util.Print(result)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
