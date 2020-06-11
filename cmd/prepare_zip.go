package cmd

import (
	"github.com/moor84/shamestock/zip"
	"github.com/spf13/cobra"
)

// zipCmd represents the zip command
var zipCmd = &cobra.Command{
	Use:   "zip path-to-directory",
	Short: "Create zip archives and stores them in the batches/batch1/zip directory",
	Long: `Create zip archives required, for example, for Adobe Stock, and store them in
the batches/batch1/zip directory.
For example, for each 123.eps and 123.jpg it will create 123.zip with both files.
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var path = args[0]
		zip.CreateZipFiles(path)
	},
}

func init() {
	prepareCmd.AddCommand(zipCmd)
}
