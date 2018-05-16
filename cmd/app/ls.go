package app

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type lsOptions struct {
	cluster   string
	namespace string
}

func newBaseLsCmd(alauda client.APIClient) *cobra.Command {
	var opts lsOptions

	lsCmd := &cobra.Command{
		Short: "List applications",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLs(alauda, &opts)
		},
	}

	lsCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")
	lsCmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", "Namespace")

	return lsCmd
}

// NewAppsCmd creates a new alauda apps command, which is a shortcut to the app ls command.
func NewAppsCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "apps"
	return cmd
}

func newLsCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "ls"
	return cmd
}

func doLs(alauda client.APIClient, opts *lsOptions) error {
	fmt.Println("[alauda] Listing applications")

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	namespace, err := util.ConfigNamespace(opts.namespace)
	if err != nil {
		return err
	}

	params := client.ListAppsParams{
		Cluster:   cluster,
		Namespace: namespace,
	}

	result, err := alauda.ListApps(&params)
	if err != nil {
		return err
	}

	printLsResult(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printLsResult(result *client.ListAppsResult) {
	header := buildLsTableHeader()
	content := buildLsTableContent(result)

	util.PrintTable(header, content)
}

func buildLsTableHeader() []string {
	return []string{"NAME", "CLUSTER", "NAMESPACE", "STATE", "IMAGE", "SIZE", "DESCRIPTION"}
}

func buildLsTableContent(result *client.ListAppsResult) [][]string {
	var content [][]string

	for _, app := range result.Apps {
		content = append(content, []string{app.Resource.Name, app.Cluster.Name, app.Namespace.Name, app.Resource.State, "", "", app.Resource.Description})

		for _, component := range app.Components {
			containers := component.Resource.Containers
			image := ""
			size := ""

			if len(containers) >= 1 {
				image = containers[0].Image
				size = fmt.Sprintf("CPU: %s, Memory: %s", containers[0].Size.CPU, containers[0].Size.Memory)
				containers = containers[1:]
			}

			content = append(content, []string{fmt.Sprintf("|-%s", component.Resource.Name), "", "",
				fmt.Sprintf("%d/%d", component.Resource.Instances.Current, component.Resource.Instances.Desired), image, size, ""})

			for _, container := range containers {
				size = fmt.Sprintf("CPU: %s, Memory: %s", container.Size.CPU, container.Size.Memory)
				content = append(content, []string{"|", "", "", "", container.Image, size, ""})
			}
		}
	}

	return content
}
