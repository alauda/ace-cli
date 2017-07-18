package node

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type inspectOptions struct {
	cluster string
}

func newInspectCmd(alauda client.APIClient) *cobra.Command {
	var opts inspectOptions

	inspectCmd := &cobra.Command{
		Use:   "inspect IP",
		Short: "Inspect a node",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("node inspect expects IP")
			}
			return doInspect(alauda, args[0], &opts)
		},
	}

	inspectCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")

	return inspectCmd
}

func doInspect(alauda client.APIClient, name string, opts *inspectOptions) error {
	fmt.Println("[alauda] Inspecting", name)

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	result, err := alauda.InspectNode(name, cluster)
	if err != nil {
		return err
	}

	err = util.Print(result)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
