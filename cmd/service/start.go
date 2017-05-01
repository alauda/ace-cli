package service

import (
	"errors"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewStartCmd creates a new start service command.
func NewStartCmd(alauda client.APIClient) *cobra.Command {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start service",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("no service name specified")
			}
			return doStart(alauda, args[0])
		},
	}

	return startCmd
}

func doStart(alauda client.APIClient, name string) error {
	util.InitializeClient(alauda)

	err := alauda.StartService(name)
	if err != nil {
		return err
	}

	return nil
}
