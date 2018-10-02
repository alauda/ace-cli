package app

import (
	"encoding/json"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type inspectOptions struct {
	cluster   string
	namespace string
}

func newInspectCmd(alauda client.APIClient) *cobra.Command {
	var opts inspectOptions

	inspectCmd := &cobra.Command{
		Use:   "inspect NAME",
		Short: "Inspect an application",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doInspect(alauda, &opts, args[0])
		},
	}

	inspectCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")
	inspectCmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", "Namespace")

	return inspectCmd
}

func doInspect(alauda client.APIClient, opts *inspectOptions, name string) error {
	fmt.Println("[alauda] Inspecting", name)

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	namespace, err := util.ConfigNamespace(opts.namespace)
	if err != nil {
		return err
	}

	result, err := alauda.InspectApp(cluster, namespace, name)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	fmt.Println("[alauda] OK")

	return nil
}
