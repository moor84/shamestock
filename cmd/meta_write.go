package cmd

import (
	"strings"

	"github.com/moor84/shamestock/meta"
	"github.com/spf13/cobra"
)

var title string
var description string
var keywords string

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write [jpeg filename]",
	Short: "Write meta attributes to a JPEG file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]

		attrs := meta.Attrs{}
		if title != "" {
			attrs.Title = &title
		} else {
			attrs.Title = nil
		}
		if description != "" {
			attrs.Description = &description
		} else {
			attrs.Description = nil
		}
		if keywords != "" {
			attrs.Keywords = strings.Split(keywords, ",")
		} else {
			attrs.Keywords = nil
		}
		meta.WriteAttrs(filename, attrs)
	},
}

func init() {
	metaCmd.AddCommand(writeCmd)

	writeCmd.Flags().StringVarP(&title, "title", "t", "", "Title")
	writeCmd.Flags().StringVarP(&description, "description", "d", "", "Description")
	writeCmd.Flags().StringVarP(&keywords, "keywords", "k", "", "Keywords (comma-separated)")
}
