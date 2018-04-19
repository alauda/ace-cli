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
	all     bool
}

func newPsCmd(alauda client.APIClient) *cobra.Command {
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
	psCmd.Flags().BoolVarP(&opts.all, "all", "a", false, "List services that belong to applications as well")

	return psCmd
}

func doPs(alauda client.APIClient, opts *psOptions) error {
	fmt.Println("[alauda] Listing services")

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
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

	var appResult *client.ListAppsResult

	if opts.all {
		appParams := client.ListAppsParams{
			Cluster: cluster,
		}

		appResult, err = alauda.ListApps(&appParams)
		if err != nil {
			return err
		}
	}

	printPsResult(result, appResult)

	fmt.Println("[alauda] OK")

	return nil
}

func printPsResult(result *client.ListServicesResult, appResult *client.ListAppsResult) {
	services := result.Services

	if appResult != nil {
		for _, app := range appResult.Apps {
			for _, service := range app.Services {
				service.Name = fmt.Sprintf("%s.%s", app.Name, service.Name)
				services = append(services, service)
			}
		}
	}

	PrintServices(services)
}

// PrintServices prints the service list in a table.
func PrintServices(services []client.Service) {
	header := buildPsTableHeader()
	content := buildPsTableContent(services)

	util.PrintTable(header, content)
}

func buildPsTableHeader() []string {
	return []string{"NAME", "IMAGE", "COMMAND", "CREATED", "SIZE", "PORTS", "COUNT", "STATE"}
}

func buildPsTableContent(services []client.Service) [][]string {
	var content [][]string

	for _, service := range services {
		image := fmt.Sprintf("%s:%s", service.ImageName, service.ImageTag)
		cpu := strconv.FormatFloat(service.Size.CPU, 'f', -1, 64)
		size := fmt.Sprintf("CPU: %s, Memory: %d", cpu, service.Size.Memory)
		ports := fmt.Sprint(service.Ports)
		count := fmt.Sprintf("%d/%d", service.HealthyInstances, service.TargetInstances)
		content = append(content, []string{service.Name, image, service.Command, service.CreatedAt, size, ports, count, service.State})
	}

	return content
}
