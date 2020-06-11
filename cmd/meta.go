package cmd

import (
	"github.com/spf13/cobra"
)

// metaCmd represents the meta command
var metaCmd = &cobra.Command{
	Use:   "meta",
	Short: "Commands for manipulating titles and keywords",
	Long: `You can scrape meta attributes from Shutterstock and/or write them into IPTC tags of the JPEGs.
shamestock meta scrape/scrapecsv allow to scrape attributes from Shutterstock web pages.
shamestock meta write/csv allow to store meta attributes in JPEG tags (that many stocks will recognize).
`,
}

func init() {
	rootCmd.AddCommand(metaCmd)
}
