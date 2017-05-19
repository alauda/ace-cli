package util

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// PrintTable prints a table with the provided header and content.
func PrintTable(header []string, content [][]string) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(header)
	table.AppendBulk(content)

	table.SetBorder(false)

	table.Render()
}

// Print marshals the object into an indented JSON and prints it.
func Print(v interface{}) error {
	marshalled, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(marshalled))

	return nil
}
