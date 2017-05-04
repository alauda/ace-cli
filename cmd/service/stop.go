package service

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewStopCmd creates a new stop service command.
func NewStopCmd(alauda client.APIClient) *cobra.Command {
	stopCmd := &cobra.Command{
		Use:   "stop NAME",
		Short: "Stop a service",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("service stop expects NAME")
			}
			return doStop(alauda, args[0])
		},
	}

	return stopCmd
}

func doStop(alauda client.APIClient, name string) error {
	fmt.Println("[alauda] Stopping", name)

	util.InitializeClient(alauda)

	err := alauda.StopService(name)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
