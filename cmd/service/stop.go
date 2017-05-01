package service

import (
	"errors"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewStopCmd creates a new stop service command.
func NewStopCmd(alauda client.APIClient) *cobra.Command {
	stopCmd := &cobra.Command{
		Use:   "stop <service>",
		Short: "Stop service",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("no service name specified")
			}
			return doStop(alauda, args[0])
		},
	}

	return stopCmd
}

func doStop(alauda client.APIClient, name string) error {
	util.InitializeClient(alauda)

	err := alauda.StopService(name)
	if err != nil {
		return err
	}

	return nil
}
