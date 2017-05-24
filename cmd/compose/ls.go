package compose

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type lsOptions struct {
	cluster string
}

// NewLsCmd creates a new compose ls command.
func NewLsCmd(alauda client.APIClient) *cobra.Command {
	var opts lsOptions

	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "List apps",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLs(alauda, &opts)
		},
	}

	lsCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")

	return lsCmd
}

func doLs(alauda client.APIClient, opts *lsOptions) error {
	fmt.Println("[alauda] Listing apps")

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	params := client.ListAppsParams{
		Cluster: cluster,
	}

	result, err := alauda.ListApps(&params)
	if err != nil {
		return err
	}

	printLsResult(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printLsResult(apps *client.ListAppsResult) {
	header := buildLsTableHeader()
	content := buildLsTableContent(apps)

	util.PrintTable(header, content)
}

func buildLsTableHeader() []string {
	return []string{"NAME", "ID", "STATE", "CREATED BY"}
}

func buildLsTableContent(result *client.ListAppsResult) [][]string {
	var content [][]string

	for _, app := range result.Apps {
		content = append(content, []string{app.Name, app.ID, app.State, app.CreatedBy})

		for _, service := range app.Services {
			content = append(content, []string{fmt.Sprintf("|-%s", service.Name), "", fmt.Sprintf("|-%s", service.State), ""})
		}
	}

	return content
}
