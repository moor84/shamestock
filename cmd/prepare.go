package cmd

import (
	"github.com/spf13/cobra"
)

// prepareCmd represents the prepare command
var prepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "Utils to prepare a batch for upload",
}

func init() {
	rootCmd.AddCommand(prepareCmd)
}
