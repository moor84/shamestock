package cmd

import (
	"github.com/moor84/shamestock/preview"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate [directory]",
	Short: "Generate JPEG previews for all .eps files in a directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var path = args[0]
		preview.CreatePreviews(path)
	},
}

func init() {
	previewCmd.AddCommand(generateCmd)
}
