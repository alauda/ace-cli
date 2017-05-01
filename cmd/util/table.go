package util

import (
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
