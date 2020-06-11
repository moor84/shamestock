package cmd

import (
	"github.com/moor84/shamestock/preview"
	"github.com/spf13/cobra"
)

var big bool

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [path-to-file]",
	Short: "Generate a JPEG preview for a vector file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var path = args[0]
		if big {
			preview.CreateBigPreview(path)
		} else {
			preview.CreateNormalPreview(path)
		}
	},
}

func init() {
	previewCmd.AddCommand(createCmd)
	createCmd.Flags().BoolVarP(&big, "big", "b", false, "Create a preview with bigger DPI")
}
