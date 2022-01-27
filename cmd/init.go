package cmd

import (
	"fmt"

	"github.com/jack5341/schnell-setup/pkg"
	"github.com/spf13/cobra"
)

// Variables for nginx configs
var domainName string
var proxyPass string

var isCompose bool

// Variables for docker-compose.yml

var path string
var imageN string
var port string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if domainName == "" && proxyPass == "" {
			fmt.Println("domain name and proxypass is required")
			return
		}

		if isCompose {
			if imageN == "" && port == "" {
				fmt.Println("imagename and port is required")
				return
			}

			pkg.GenerateNginxConf(domainName, proxyPass, path)
			pkg.GenerateCompose(path, imageN, port)
			return
		}

		pkg.GenerateNginxConf(domainName, proxyPass, path)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVar(&domainName, "domain", "", "domain name for initialize nginx.conf file")
	initCmd.Flags().StringVar(&proxyPass, "proxy", "", "proxy_pass for initialize nginx.conf file")
	initCmd.Flags().StringVar(&path, "path", "", "path for initialize nginx.conf file")
	initCmd.Flags().StringVar(&imageN, "imagename", "", "image name for docker-compose.yml file")
	initCmd.Flags().StringVar(&port, "port", "", "port for docker-compose.yml file")
	initCmd.Flags().BoolVar(&isCompose, "compose", false, "initialize docker-compose.yml file")
}
