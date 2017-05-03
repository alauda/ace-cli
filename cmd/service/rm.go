package service

import (
	"errors"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewRmCmd creates a new remove service command.
func NewRmCmd(alauda client.APIClient) *cobra.Command {
	startCmd := &cobra.Command{
		Use:   "rm NAME",
		Short: "Remove service",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("service rm expects NAME")
			}
			return doRm(alauda, args[0])
		},
	}

	return startCmd
}

func doRm(alauda client.APIClient, name string) error {
	util.InitializeClient(alauda)

	err := alauda.RemoveService(name)
	if err != nil {
		return err
	}

	return nil
}
