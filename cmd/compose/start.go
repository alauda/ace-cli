package compose

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewStartCmd creates a new compose start command.
func NewStartCmd(alauda client.APIClient) *cobra.Command {
	startCmd := &cobra.Command{
		Use:   "start NAME",
		Short: "Start an application",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("compose start expects NAME")
			}
			return doStart(alauda, args[0])
		},
	}

	return startCmd
}

func doStart(alauda client.APIClient, name string) error {
	fmt.Println("[alauda] Starting", name)

	util.InitializeClient(alauda)

	err := alauda.StartApp(name)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
