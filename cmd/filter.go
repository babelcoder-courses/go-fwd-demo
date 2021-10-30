/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"fwd-search-api/internal/engine"

	"github.com/spf13/cobra"
)

// filterCmd represents the filter command
var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stderr := cmd.ErrOrStderr()
		stdout := cmd.OutOrStdout()
		term := args[0]
		file, err := cmd.Flags().GetString("file-path")
		if err != nil || file == "" {
			fmt.Fprintln(stderr, "invalid file path")
		}

		e := engine.NewFileEngine("", file)
		rs, err := e.Filter(term)
		if err != nil {
			fmt.Fprintln(stderr, err)
		}

		for _, r := range rs {
			fmt.Fprintf(stdout, "%d: %s\n", r.ID, r.Title)
		}
	},
}

func init() {
	RootCmd.AddCommand(filterCmd)
	filterCmd.Flags().StringP("file-path", "f", "", "The saved path")
}
