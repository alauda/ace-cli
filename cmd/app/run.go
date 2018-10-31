package app

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type runOptions struct {
	cluster   string
	namespace string
}

// NewRunCmd creates a new alauda app run command.
func NewRunCmd(alauda client.APIClient) *cobra.Command {
	var opts runOptions

	runCmd := &cobra.Command{
		Use:   "run NAME IMAGE",
		Short: "Run an application",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("app run expects NAME and IMAGE")
			}
			return doRun(alauda, args[0], args[1], &opts)
		},
	}

	runCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")
	runCmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", "Namespace")

	return runCmd
}

func doRun(alauda client.APIClient, name string, image string, opts *runOptions) error {
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

	err = alauda.RunApp(cluster, namespace, name, image)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
