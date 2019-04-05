package namespace

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

	lsCmd := &cobra.Command{
		Use:   "inspect NAME",
		Short: "Inspect namespace",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("namespace inspect expects NAME")
			}
			return doInspect(alauda, &opts, args[0])
		},
	}

	lsCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")

	return lsCmd
}

func doInspect(alauda client.APIClient, opts *inspectOptions, name string) error {
	fmt.Println("[alauda] Inspecting", name)

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	result, err := alauda.InspectNamespace(cluster, name)
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
