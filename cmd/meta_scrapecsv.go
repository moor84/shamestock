package cmd

import (
	"github.com/moor84/shamestock/meta"
	"github.com/spf13/cobra"
)

// scrapecsvCmd represents the scrape command
var scrapecsvCmd = &cobra.Command{
	Use:   "scrapecsv [path to CSV file]",
	Short: "Scrape meta attributes from Shuttertock using CSV files.",
	Long: `Input file example:

1234.eps,https://www.shutterstock.com/image-vector/christmas-banners-1234
2345.eps,https://www.shutterstock.com/image-vector/something-2345

Will create an output file attrs.csv (in the same directory as an input file):

1234.eps,Christmas Banner,"web,bannner,holiday",https://www.shutterstock.com/image-vector/christmas-banners-1234
2345.eps,Something Else,"keywords,here",https://www.shutterstock.com/image-vector/something-2345
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var path = args[0]
		meta.ScrapeAttrsCSV(path)
	},
}

func init() {
	metaCmd.AddCommand(scrapecsvCmd)
}
