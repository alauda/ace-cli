package space

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewLsCmd creates a new space ls command.
func NewLsCmd(alauda client.APIClient) *cobra.Command {
	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "List spaces",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLs(alauda)
		},
	}
	return lsCmd
}

func doLs(alauda client.APIClient) error {
	fmt.Println("[alauda] Listing spaces")

	util.InitializeClient(alauda)

	result, err := alauda.ListSpaces()
	if err != nil {
		return err
	}

	printLsResult(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printLsResult(spaces *client.ListSpacesResult) {
	header := buildLsTableHeader()
	content := buildLsTableContent(spaces)

	util.PrintTable(header, content)
}

func buildLsTableHeader() []string {
	return []string{"NAME", "CREATED", "STATUS"}
}

func buildLsTableContent(result *client.ListSpacesResult) [][]string {
	var content [][]string

	for i := 0; i < len(result.Spaces); i++ {
		space := result.Spaces[i]
		content = append(content, []string{space.Name, space.CreatedAt, space.Status})
	}

	return content
}
