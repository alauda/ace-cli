package volume

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewInspectCmd creates a new volume inspect command.
func NewInspectCmd(alauda client.APIClient) *cobra.Command {
	inspectCmd := &cobra.Command{
		Use:   "inspect NAME",
		Short: "Inspect a volume",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("volume inspect expects NAME")
			}
			return doInspect(alauda, args[0])
		},
	}

	return inspectCmd
}

func doInspect(alauda client.APIClient, name string) error {
	fmt.Println("[alauda] Inspecting", name)

	util.InitializeClient(alauda)

	id, err := getVolumeID(alauda, name)
	if err != nil {
		return err
	}

	result, err := alauda.InspectVolume(id)
	if err != nil {
		return err
	}

	err = printVolume(result)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}

func printVolume(lb *client.Volume) error {
	marshalled, err := json.MarshalIndent(lb, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(marshalled))

	return nil
}
