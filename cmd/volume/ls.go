package volume

import (
	"fmt"
	"strconv"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type lsOptions struct {
	cluster string
}

func newBaseLsCmd(alauda client.APIClient) *cobra.Command {
	var opts lsOptions

	lsCmd := &cobra.Command{
		Short: "List volumes",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLs(alauda, &opts)
		},
	}

	lsCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")

	return lsCmd
}

// NewVolumesCmd creates a new alauda volumes command, which is a shortcut to the volume ls command.
func NewVolumesCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "volumes"
	return cmd
}

func newLsCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "ls"
	return cmd
}

func doLs(alauda client.APIClient, opts *lsOptions) error {
	fmt.Println("[alauda] Listing volumes")

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	clusterID, err := getClusterID(alauda, cluster)
	if err != nil {
		return err
	}

	params := client.ListVolumesParams{
		ClusterID: clusterID,
	}

	result, err := alauda.ListVolumes(&params)
	if err != nil {
		return err
	}

	printLsResult(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printLsResult(result *client.ListVolumesResult) {
	header := buildLsTableHeader()
	content := buildLsTableContent(result)

	util.PrintTable(header, content)
}

func buildLsTableHeader() []string {
	return []string{"NAME", "DRIVER", "STATE", "SIZE", "CREATED AT", "CREATED BY"}
}

func buildLsTableContent(result *client.ListVolumesResult) [][]string {
	var content [][]string

	for _, volume := range result.Volumes {
		content = append(content, []string{volume.Name, volume.Driver, volume.State,
			strconv.Itoa(volume.Size), volume.CreatedAt, volume.CreatedBy})
	}

	return content
}
