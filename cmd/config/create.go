package config

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type createOptions struct {
	description string
	items       []string
}

func newCreateCmd(alauda client.APIClient) *cobra.Command {
	var opts createOptions

	createCmd := &cobra.Command{
		Use:   "create NAME",
		Short: "Create a new configuration",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("config create expects NAME")
			}
			return doCreate(alauda, args[0], &opts)
		},
	}

	createCmd.Flags().StringVarP(&opts.description, "description", "d", "", "Description")
	createCmd.Flags().StringSliceVarP(&opts.items, "item", "i", []string{}, "Configuration items")

	return createCmd
}

func doCreate(alauda client.APIClient, name string, opts *createOptions) error {
	fmt.Println("[alauda] Creating", name)

	util.InitializeClient(alauda)

	items, err := util.ParseKeyValues(opts.items)
	if err != nil {
		return err
	}

	data := client.CreateConfigData{
		Name:        name,
		Description: opts.description,
	}

	data.Content = make([]client.ConfigItem, len(items))

	i := 0
	for k, v := range items {
		item := client.ConfigItem{
			Key:   k,
			Value: v,
		}
		data.Content[i] = item
		i++
	}

	err = alauda.CreateConfig(&data)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
