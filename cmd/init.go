package cmd

import (
	"fmt"

	"github.com/jack5341/schnell-setup/pkg"
	"github.com/spf13/cobra"
)

var domainName string
var proxyPass string
var path string
var isCompose bool

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if domainName == "" && proxyPass == "" {
			fmt.Println("domain name and proxypass is required")
			return
		}

		pkg.GenerateNginxConf(domainName, proxyPass, path)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().StringVar(&domainName, "domain", "", "domain name for initialize nginx.conf file")
	initCmd.PersistentFlags().StringVar(&proxyPass, "proxy", "", "proxy_pass for initialize nginx.conf file")
	initCmd.PersistentFlags().StringVar(&path, "path", "", "path for initialize nginx.conf file")
	initCmd.Flags().BoolVar(&isCompose, "compose", false, "initialize docker-compose.yml file")
}
