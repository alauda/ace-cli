package service

import (
	"fmt"
	"strconv"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type psOptions struct {
	cluster string
}

// NewPsCmd creates a new service ps command.
func NewPsCmd(alauda client.APIClient) *cobra.Command {
	var opts psOptions

	psCmd := &cobra.Command{
		Use:   "ps",
		Short: "List services",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doPs(alauda, &opts)
		},
	}

	psCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")

	return psCmd
}

func doPs(alauda client.APIClient, opts *psOptions) error {
	util.InitializeClient(alauda)

	cluster, err := configCluster(opts.cluster)
	if err != nil {
		return err
	}

	params := client.ListServicesParams{
		Cluster: cluster,
	}

	result, err := alauda.ListServices(&params)
	if err != nil {
		return err
	}

	printPsResult(result)

	return nil
}

func printPsResult(services *client.ListServicesResult) {
	header := buildPsTableHeader()
	content := buildPsTableContent(services)

	util.PrintTable(header, content)
}

func buildPsTableHeader() []string {
	return []string{"NAME", "IMAGE", "COMMAND", "CREATED", "SIZE", "PORTS", "COUNT", "STATUS"}
}

func buildPsTableContent(services *client.ListServicesResult) [][]string {
	var content [][]string

	for i := 0; i < services.Count; i++ {
		service := services.Results[i]
		image := fmt.Sprintf("%s:%s", service.ImageName, service.ImageTag)
		cpu := strconv.FormatFloat(service.Size.CPU, 'f', -1, 64)
		size := fmt.Sprintf("CPU: %s, Memory: %d", cpu, service.Size.Memory)
		ports := fmt.Sprint(service.Ports)
		count := fmt.Sprintf("%d/%d", service.HealthyInstances, service.TargetInstances)
		content = append(content, []string{service.Name, image, service.Command, service.Created, size, ports, count, service.Status})
	}

	return content
}
