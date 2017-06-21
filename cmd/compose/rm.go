package compose

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewRmCmd creates a new compose rm command.
func NewRmCmd(alauda client.APIClient) *cobra.Command {
	startCmd := &cobra.Command{
		Use:   "rm NAME",
		Short: "Remove an application",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("compose rm expects NAME")
			}
			return doRm(alauda, args[0])
		},
	}

	return startCmd
}

func doRm(alauda client.APIClient, name string) error {
	fmt.Println("[alauda] Removing", name)

	util.InitializeClient(alauda)

	err := alauda.RemoveApp(name)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
