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
	"fwd-search-api/internal/engine"
	"log"

	"github.com/spf13/cobra"
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "dump content from the internet into specific file or database",
	Long:  "This command will download content from the url param and then save it into file-path or database",
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")
		if err != nil || validate.Var(url, "required,url") != nil {
			log.Fatalln("invalid URL")
		}

		kind, err := cmd.Flags().GetString("type")
		if err != nil || validate.Var(kind, "oneof=file db") != nil {
			log.Fatalln("invalid type")
		}

		var e engine.Engine
		if kind == "file" {
			file, err := cmd.Flags().GetString("file-path")
			if err != nil || file == "" {
				log.Fatalln("invalid file path")
			}

			e = engine.NewFileEngine(url, file)
		} else {
			e = engine.NewDBEngine(url)
		}

		if err := e.Dump(); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(dumpCmd)

	dumpCmd.Flags().StringP("url", "u", "", "The API URL")
	dumpCmd.Flags().StringP("file-path", "f", "", "The saved path")
	dumpCmd.Flags().StringP("type", "t", "file", "Dump content into file or database")
}
