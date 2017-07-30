package space

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

func newBaseLsCmd(alauda client.APIClient) *cobra.Command {
	lsCmd := &cobra.Command{
		Short: "List spaces",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLs(alauda)
		},
	}
	return lsCmd
}

// NewSpacesCmd creates a new alauda spaces command, which is a shortcut to the space ls command.
func NewSpacesCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "spaces"
	return cmd
}

func newLsCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "ls"
	return cmd
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

func printLsResult(result *client.ListSpacesResult) {
	header := buildLsTableHeader()
	content := buildLsTableContent(result)

	util.PrintTable(header, content)
}

func buildLsTableHeader() []string {
	return []string{"NAME", "CREATED", "STATE"}
}

func buildLsTableContent(result *client.ListSpacesResult) [][]string {
	var content [][]string

	for _, space := range result.Spaces {
		content = append(content, []string{space.Name, space.CreatedAt, space.State})
	}

	return content
}
