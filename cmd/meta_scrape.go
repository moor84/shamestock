package cmd

import (
	"fmt"
	"strings"

	"github.com/moor84/shamestock/meta"
	"github.com/spf13/cobra"
)

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape [url]",
	Short: "Scrape (and print) meta attributes from Shuttertock URL.",
	Long: `Example:
shamestock scrape https://www.shutterstock.com/image-vector/christmas-banners-12345
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var url = args[0]
		var attrs = meta.Scrape(url)
		fmt.Println("Title: " + *attrs.Title)
		fmt.Println("Keywords: " + strings.Join(attrs.Keywords, ", "))
	},
}

func init() {
	metaCmd.AddCommand(scrapeCmd)
}
