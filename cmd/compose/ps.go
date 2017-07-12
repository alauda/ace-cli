package compose

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/service"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

func newPsCmd(alauda client.APIClient) *cobra.Command {
	psCmd := &cobra.Command{
		Use:   "ps NAME",
		Short: "List services of the application",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("compose ps expects NAME")
			}
			return doPs(alauda, args[0])
		},
	}

	return psCmd
}

func doPs(alauda client.APIClient, name string) error {
	fmt.Println("[alauda] Listing services for", name)

	util.InitializeClient(alauda)

	result, err := alauda.InspectApp(name)
	if err != nil {
		return err
	}

	service.PrintServices(result.Services)

	fmt.Println("[alauda] OK")

	return nil
}
