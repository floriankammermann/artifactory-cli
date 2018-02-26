package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/floriankammermann/artifactory-cli/artapi"
	"github.com/floriankammermann/artifactory-cli/types"
	"text/tabwriter"
	"os"
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
		repos := make([]types.Repository,0)
		h.ExecRequest("/api/repositories/", &repos)
		// create a new tabwriter
		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
		fmt.Fprintln(w, "Key\tType\tUrl\t")
		for _, repo := range repos {
			fmt.Fprintf(w, "%s\t%s\t%s\t\n", repo.Key, repo.Type, repo.Url)
		}
		fmt.Fprintln(w)
		w.Flush()
	},
}

func init() {
	repositoryCmd.AddCommand(listCmd)
	RootCmd.AddCommand(repositoryCmd)
}
