package app

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type lsOptions struct {
	project   string
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

	lsCmd.Flags().StringVarP(&opts.project, "project", "p", "", "Project")
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

// NewLsCmd creates a new alauda app ls command.
func NewLsCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "ls"
	return cmd
}

func doLs(alauda client.APIClient, opts *lsOptions) error {
	fmt.Println("[alauda] Listing applications")

	util.InitializeClient(alauda)

	project, err := util.ConfigProject(opts.project)
	if err != nil {
		return err
	}

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	namespace, err := util.ConfigNamespace(opts.namespace)
	if err != nil {
		return err
	}

	params := client.ListAppsParams{
		Project: project,
	}

	result, err := alauda.ListApps(cluster, namespace, &params)
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
	return []string{"NAME", "NAMESPACE", "INSTANCES", "IMAGE", "SIZE"}
}

func buildLsTableContent(result *client.ListAppsResult) [][]string {
	var content [][]string

	for _, app := range *result {
		crd, err := app.ExtractAppCrd()
		if err != nil {
			return content
		}

		content = append(content, []string{crd.Name, crd.Namespace, "", "", ""})

		deployments, err := app.ExtractDeployments()
		if err != nil {
			return content
		}

		for _, deployment := range deployments {
			containers := deployment.Spec.Template.Spec.Containers
			var image string
			var size string

			if len(containers) >= 1 {
				image = containers[0].Image
				size = fmt.Sprintf("CPU: %s, Memory: %s", containers[0].Resources.Requests.Cpu().String(), containers[0].Resources.Requests.Memory().String())
				containers = containers[1:]
			}

			content = append(content, []string{fmt.Sprintf("|-%s", deployment.Name), "",
				fmt.Sprintf("%d/%d", deployment.Status.Replicas, *deployment.Spec.Replicas), image, size})

			for _, container := range containers {
				image = container.Image
				size = fmt.Sprintf("CPU: %s, Memory: %s", container.Resources.Requests.Cpu().String(), container.Resources.Requests.Memory().String())
				content = append(content, []string{"|", "", "", image, size})
			}
		}
	}

	return content
}
