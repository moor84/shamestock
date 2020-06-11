package cmd

import (
	"github.com/moor84/shamestock/meta"
	"github.com/spf13/cobra"
)

// csv represents the csv command
var csv = &cobra.Command{
	Use:   "csv [path to CSV file]",
	Short: "Write meta attributes to JPEG files from an input CSV file",
	Long: `Input file example:

1234.eps,Christmas Banner,"web,bannner,holiday",https://www.shutterstock.com/image-vector/christmas-banners-1234
2345.eps,Something Else,"keywords,here",https://www.shutterstock.com/image-vector/something-2345

This command will look for 1234.jpg and 2345.jpg in the same directory where the CSV file is located
and write the keywords and titles to theirs IPTC tags (that most of stocks will recognise).
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var path = args[0]
		meta.WriteAttrsFromCSV(path)
	},
}

func init() {
	metaCmd.AddCommand(csv)
}
