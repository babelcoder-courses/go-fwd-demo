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
	"fwd-search-api/config"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
)

var validate *validator.Validate

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "fwd-search-api",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	validate = validator.New()

	if err := config.Load(); err != nil {
		log.Fatalln(err)
	}

	if err := config.InitDB(); err != nil {
		log.Fatalln(err)
	}
}
