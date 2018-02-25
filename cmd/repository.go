package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/floriankammermann/vcloud-cli/vcdapi"
	"github.com/spf13/viper"
	"github.com/floriankammermann/artifactory-cli/artapi"
)


// repositoryCmd represents the query command
var repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "interact with repositories",
	Long: "interact with repositories",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all repositories",
	Long: "get all allocated ips of an org network",
	Run: func(cmd *cobra.Command, args []string) {
		h := artapi.InitHttpClient()
		h.ExecRequest("/api/repositories", nil)
	},
}

func init() {
	repositoryCmd.AddCommand(listCmd)
	RootCmd.AddCommand(repositoryCmd)
}
