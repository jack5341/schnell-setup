/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"github.com/jack5341/schnell-setup/pkg"
	"github.com/spf13/cobra"
)

var domainName string
var proxyPass string
var isCompose bool

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.GenerateNginxConf(domainName, proxyPass)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().StringVar(&domainName, "domain", "", "domain name for initialize nginx.conf file")
	initCmd.PersistentFlags().StringVar(&proxyPass, "proxy", "", "proxy_pass for initialize nginx.conf file")
	initCmd.Flags().BoolVar(&isCompose, "compose", false, "initialize docker-compose.yml file")
}
