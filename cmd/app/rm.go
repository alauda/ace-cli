package app

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type rmOptions struct {
	cluster   string
	namespace string
}

// NewRmCmd returns a new alauda app rm command.
func NewRmCmd(alauda client.APIClient) *cobra.Command {
	var opts rmOptions

	rmCmd := &cobra.Command{
		Use:   "rm NAME",
		Short: "Remove an application",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("app rm expects NAME")
			}
			return doRm(alauda, &opts, args[0])
		},
	}

	rmCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")
	rmCmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", "Namespace")

	return rmCmd
}

func doRm(alauda client.APIClient, opts *rmOptions, name string) error {
	fmt.Println("[alauda] Removing", name)

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	namespace, err := util.ConfigNamespace(opts.namespace)
	if err != nil {
		return err
	}

	err = alauda.RemoveApp(cluster, namespace, name)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
