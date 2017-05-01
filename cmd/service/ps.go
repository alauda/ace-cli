package service

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type psOptions struct {
	cluster string
}

// NewPsCmd creates a new service ps command.
func NewPsCmd(alauda client.AlaudaClient) *cobra.Command {
	var opts psOptions

	psCmd := &cobra.Command{
		Use:   "ps",
		Short: "List services",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doPs(alauda, opts)
		},
	}

	psCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")

	return psCmd
}

func doPs(alauda client.AlaudaClient, opts psOptions) error {
	util.InitializeClient(alauda)

	params := client.ListServicesParams{
		Cluster: opts.cluster,
	}

	result, err := alauda.ListServices(&params)
	if err != nil {
		return err
	}

	header := buildPsTableHeader()
	content := buildPsTableContent(result)

	util.PrintTable(header, content)

	return nil
}

func buildPsTableHeader() []string {
	return []string{"NAME", "IMAGE", "COMMAND", "CREATED", "SIZE", "COUNT", "STATUS"}
}

func buildPsTableContent(services *client.ListServicesResult) [][]string {
	var content [][]string

	for i := 0; i < services.Count; i++ {
		service := services.Results[i]
		image := fmt.Sprintf("%s:%s", service.ImageName, service.ImageTag)
		size := fmt.Sprintf("CPU: %d, Memory: %d", service.Size.CPU, service.Size.Memory)
		count := fmt.Sprintf("%d/%d", service.HealthyInstances, service.TargetInstances)
		content = append(content, []string{service.Name, image, service.Command, service.Created, size, count, service.Status})
	}

	return content
}
