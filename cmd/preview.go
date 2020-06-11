package cmd

import (
	"github.com/spf13/cobra"
)

// previewCmd represents the preview command
var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Generate JPEG previews for vectero files",
	Long: `You can create a preview for a single file:

shametock preview create path/to/file.eps

or for all files in a folder:

shametock preview generate path/to/directory/
`,
}

func init() {
	rootCmd.AddCommand(previewCmd)
}
