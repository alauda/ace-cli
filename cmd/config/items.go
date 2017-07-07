package config

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewItemsCmd creates a new config items command.
func NewItemsCmd(alauda client.APIClient) *cobra.Command {
	itemsCmd := &cobra.Command{
		Use:   "items NAME",
		Short: "List config items of a configuration",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("config items expects NAME")
			}
			return doItems(alauda, args[0])
		},
	}

	return itemsCmd
}

func doItems(alauda client.APIClient, name string) error {
	fmt.Println("[alauda] Listing config items for", name)

	util.InitializeClient(alauda)

	result, err := alauda.InspectConfig(name)
	if err != nil {
		return err
	}

	printItems(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printItems(config *client.Config) {
	header := buildItemsTableHeader()
	content := buildItemsTableContent(config)

	util.PrintTable(header, content)
}

func buildItemsTableHeader() []string {
	return []string{"KEY", "VALUE"}
}

func buildItemsTableContent(config *client.Config) [][]string {
	var content [][]string

	for _, item := range config.Content {
		content = append(content, []string{item.Key, item.Value})
	}

	return content
}
