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
	"github.com/spf13/cobra"

	"github.com/goakshit/sauron/cmd/newcmd"
	"github.com/goakshit/sauron/cmd/paybackcmd"
	"github.com/goakshit/sauron/cmd/reportcmd"
	"github.com/goakshit/sauron/cmd/updatecmd"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sauron",
	Short: "Sauron's eye sees everything",
	Long:  "",
	// Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {

	rootCmd.AddCommand(
		newcmd.GetNewCmd(),
		updatecmd.GetUpdateCmd(),
		paybackcmd.GetPaybackCmd(),
		reportcmd.GetReportCmd(),
	)
}
