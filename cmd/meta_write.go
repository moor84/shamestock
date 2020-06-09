/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"strings"

	"github.com/moor84/shamestock/meta"
	"github.com/spf13/cobra"
)

var Title string
var Description string
var Keywords string

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write [jpeg filename]",
	Short: "Write meta attributes to a jpeg file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]

		attrs := meta.Attrs{}
		if Title != "" {
			attrs.Title = &Title
		} else {
			attrs.Title = nil
		}
		if Description != "" {
			attrs.Description = &Description
		} else {
			attrs.Description = nil
		}
		if Keywords != "" {
			attrs.Keywords = strings.Split(Keywords, ",")
		} else {
			attrs.Keywords = nil
		}
		meta.WriteAttrs(filename, attrs)
	},
}

func init() {
	metaCmd.AddCommand(writeCmd)

	writeCmd.Flags().StringVarP(&Title, "title", "t", "", "Title")
	writeCmd.Flags().StringVarP(&Description, "description", "d", "", "Description")
	writeCmd.Flags().StringVarP(&Keywords, "keywords", "k", "", "Keywords (comma-separated)")
}
