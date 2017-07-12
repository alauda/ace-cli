package config

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

func newRmCmd(alauda client.APIClient) *cobra.Command {
	rmCmd := &cobra.Command{
		Use:   "rm NAME",
		Short: "Remove a configuration",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("config rm expects NAME")
			}
			return doRm(alauda, args[0])
		},
	}

	return rmCmd
}

func doRm(alauda client.APIClient, name string) error {
	fmt.Println("[alauda] Removing", name)

	util.InitializeClient(alauda)

	err := alauda.RemoveConfig(name)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
