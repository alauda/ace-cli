package compose

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

func newInspectCmd(alauda client.APIClient) *cobra.Command {
	inspectCmd := &cobra.Command{
		Use:   "inspect NAME",
		Short: "Inspect an application",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("compose inspect expects NAME")
			}
			return doInspect(alauda, args[0])
		},
	}

	return inspectCmd
}

func doInspect(alauda client.APIClient, name string) error {
	fmt.Println("[alauda] Inspecting", name)

	util.InitializeClient(alauda)

	result, err := alauda.InspectApp(name)
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
