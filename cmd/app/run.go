package app

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type runOptions struct {
	cluster   string
	namespace string
}

func newRunCmd(alauda client.APIClient) *cobra.Command {
	var opts runOptions

	runCmd := &cobra.Command{
		Use:   "run NAME",
		Short: "Run an application",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doRun(alauda, &opts, args[0])
		},
	}

	runCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")
	runCmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", "Namespace")

	return runCmd
}

func doRun(alauda client.APIClient, opts *runOptions, name string) error {
	fmt.Println("[alauda] Running", name)

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	namespace, err := util.ConfigNamespace(opts.namespace)
	if err != nil {
		return err
	}

	err = alauda.RunApp(cluster, namespace, name)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
