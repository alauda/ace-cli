package registry

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

func newProjectsCmd(alauda client.APIClient) *cobra.Command {
	projectsCmd := &cobra.Command{
		Use:   "projects NAME",
		Short: "List projects of a registry",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("registry projects expects NAME")
			}
			return doProjects(alauda, args[0])
		},
	}

	return projectsCmd
}

func doProjects(alauda client.APIClient, registryName string) error {
	fmt.Println("[alauda] Listing projects in", registryName)

	util.InitializeClient(alauda)

	result, err := alauda.ListRegistryProjects(registryName)
	if err != nil {
		return err
	}

	printProjectsResult(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printProjectsResult(result *client.ListRegistryProjectsResult) {
	header := buildProjectsTableHeader()
	content := buildProjectsTableContent(result)

	util.PrintTable(header, content)
}

func buildProjectsTableHeader() []string {
	return []string{"NAME", "ID", "CREATED BY", "REPO COUNT"}
}

func buildProjectsTableContent(result *client.ListRegistryProjectsResult) [][]string {
	var content [][]string

	for _, project := range result.Projects {
		content = append(content, []string{project.Name, project.ID, project.CreatedBy, strconv.Itoa(project.RepoCount)})
	}

	return content
}
