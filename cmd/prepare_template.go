package cmd

import (
	"github.com/moor84/shamestock/meta"
	"github.com/spf13/cobra"
)

// templateCmd represents the zip command
var templateCmd = &cobra.Command{
	Use:   "template path-to-directory",
	Short: "Creates a template for a CSV file, batches/batch1/urls.csv",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var path = args[0]
		meta.CreateCSVtemplate(path)
	},
}

func init() {
	prepareCmd.AddCommand(templateCmd)
}
