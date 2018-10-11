package app

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type deleteOptions struct {
	cluster   string
	namespace string
}

func newDeleteCmd(alauda client.APIClient) *cobra.Command {
	var opts deleteOptions

	deleteCmd := &cobra.Command{
		Use:   "delete NAME",
		Short: "Delete an application",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doDelete(alauda, &opts, args[0])
		},
	}

	deleteCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")
	deleteCmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", "Namespace")

	return deleteCmd
}

func doDelete(alauda client.APIClient, opts *deleteOptions, name string) error {
	fmt.Println("[alauda] Deleting", name)

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	namespace, err := util.ConfigNamespace(opts.namespace)
	if err != nil {
		return err
	}

	err = alauda.DeleteApp(cluster, namespace, name)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
