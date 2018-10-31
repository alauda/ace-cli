package app

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type yamlOptions struct {
	cluster   string
	namespace string
}

// NewYamlCmd creates a new alauda app yaml command.
func NewYamlCmd(alauda client.APIClient) *cobra.Command {
	var opts yamlOptions

	yamlCmd := &cobra.Command{
		Use:   "yaml NAME",
		Short: "Retrieve the YAML of an application",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doGetYaml(alauda, &opts, args[0])
		},
	}

	yamlCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")
	yamlCmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", "Namespace")

	return yamlCmd
}

func doGetYaml(alauda client.APIClient, opts *yamlOptions, name string) error {
	fmt.Println("[alauda] Getting YAML for", name)

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	namespace, err := util.ConfigNamespace(opts.namespace)
	if err != nil {
		return err
	}

	result, err := alauda.GetAppYaml(cluster, namespace, name)
	if err != nil {
		return err
	}

	fmt.Println(result)

	fmt.Println("[alauda] OK")

	return nil
}
