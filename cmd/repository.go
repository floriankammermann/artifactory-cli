package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/floriankammermann/artifactory-cli/artapi"
	"github.com/floriankammermann/artifactory-cli/types"
	"text/tabwriter"
	"os"
	"github.com/spf13/viper"
	"bytes"
	"encoding/json"
)

var repoType string
var repoPackageType string

// repositoryCmd represents the query command
var repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "interact with repositories",
	Long: "interact with repositories",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all repositories",
	Long: "list all available repositories",
	Run: func(cmd *cobra.Command, args []string) {
		h := artapi.InitHttpClient()
		repos := make([]types.Repository,0)
		h.ExecRequest("/api/repositories/", "GET", nil, &repos)
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

var existsCmd = &cobra.Command{
	Use:   "exists",
	Short: "check if repository exists",
	Long: "check if the repository with the given name exists",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("no repository specified")
			os.Exit(1)
		}
		if len(args) > 1 {
			fmt.Println("can only check for one repo")
			os.Exit(1)
		}
		if viper.GetString("verbose") == "true" {
			fmt.Printf("check for existence of repo: %s\n", args[0])
		}
		h := artapi.InitHttpClient()
		repos := make([]types.Repository,0)
		h.ExecRequest("/api/repositories/", "GET", nil, &repos)
		for _, repo := range repos {
			if repo.Key == args[0] {
				fmt.Printf("found repository with key: %s\n", repo.Key)
				os.Exit(0)
			}
		}
		fmt.Printf("repository not found with key: %s\n", args[0])
	},
}

func getRepoPackageTypes() []string {
	repoPackageTypes := []string{"maven","gradle","ivy","sbt","nuget","gems","npm","bower","debian","composer","pypi",
								 "docker","vagrant","gitlfs","yum","conan","chef","puppet","generic"}
	return repoPackageTypes
}

func getRepoPackageTypesAsString() string {
	var buffer bytes.Buffer
	for _, repoPackageType := range getRepoPackageTypes() {
		buffer.WriteString(repoPackageType)
		buffer.WriteString(", ")
	}
	return buffer.String()
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func printTextAndExit(text string) {
	fmt.Println(text)
	os.Exit(1)
}

func getRepoLayout(repoType string) string {
	repoLayoutDef := make(map[string]string)
	repoLayoutDef["bower"] = "bower-default"
	repoLayoutDef["composer"] = "composer-default"
	repoLayoutDef["conan"] = "conan-default"
	repoLayoutDef["gradle"] = "gradle-default"
	repoLayoutDef["ivy"] = "ivy-default"
	repoLayoutDef["maven"] = "maven-2-default"
	repoLayoutDef["npm"] = "npm-default"
	repoLayoutDef["nuget"] = "nuget-default"
	repoLayoutDef["puppet"] = "puppet-default"
	repoLayoutDef["sbt"] = "sbt-default"
	repoLayoutDef["vcs"] = "vcs-default"
	repoLayoutDef["rest"] = "simple-default"

	if repoLayoutDef[repoType] == "" {
		return "simple-default"
	}
	return repoLayoutDef[repoType]
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create repository",
	Long: "create repository with its parameters",
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetString("verbose") == "true" {
			fmt.Printf("create repo: %s, type: %s, package type: %s\n", args[0], repoType, repoPackageType)
		}
		if len(args) == 0 {
			printTextAndExit("no repository specified")
		}
		if len(args) > 1 {
			printTextAndExit("only one repo name allowed")
		}
		if repoType == "" {
			printTextAndExit("you have to provide the repoType")
		}
		if repoType != "local" {
			printTextAndExit("only local repos are supoorted right now")
		}
		if repoPackageType == "" {
			printTextAndExit("provide packageType")
		}
		if ! contains(getRepoPackageTypes(), repoPackageType) {
			printTextAndExit("provide a valid packageType")
		}
		repositoryDetails := types.CreateRepositoryDetails(repoType, repoPackageType, args[0], getRepoLayout(repoType))
		repositoryDetailsBytes, _ := json.Marshal(repositoryDetails)
		h := artapi.InitHttpClient()
		h.ExecRequest("/api/repositories/"+repositoryDetails.Key, "PUT", repositoryDetailsBytes, nil)
	},
}

func init() {
	repositoryCmd.AddCommand(listCmd)
	repositoryCmd.AddCommand(existsCmd)

	createCmd.Flags().StringVarP(&repoType, "repoType", "t", "","repository type [local]")
	createCmd.Flags().StringVarP(&repoPackageType, "packageType", "p", "","repository package type ["+getRepoPackageTypesAsString()+"]")
	repositoryCmd.AddCommand(createCmd)

	RootCmd.AddCommand(repositoryCmd)
}
