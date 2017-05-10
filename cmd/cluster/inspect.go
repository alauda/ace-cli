package cluster

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewInspectCmd creates a new start service command.
func NewInspectCmd(alauda client.APIClient) *cobra.Command {
	inspectCmd := &cobra.Command{
		Use:   "inspect NAME",
		Short: "Inspect a cluster",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("cluster inspect expects NAME")
			}
			return doInspect(alauda, args[0])
		},
	}

	return inspectCmd
}

func doInspect(alauda client.APIClient, name string) error {
	fmt.Println("[alauda] Inspecting", name)

	util.InitializeClient(alauda)

	result, err := alauda.InspectCluster(name)
	if err != nil {
		return err
	}

	err = printCluster(result)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}

func printCluster(cluster *client.Cluster) error {
	marshalled, err := json.MarshalIndent(cluster, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(marshalled))

	return nil
}
