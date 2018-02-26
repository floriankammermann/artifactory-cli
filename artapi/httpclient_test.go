package artapi

import (
	"testing"
	"os"
	"github.com/floriankammermann/artifactory-cli/types"
	"fmt"
	"encoding/json"
)

func createHttpClient() *HttpClient {
	h := InitHttpClient()
	h.user = os.Getenv("USERNAME")
	h.password = os.Getenv("PASSWORD")
	h.url = os.Getenv("URL")
	return h
}

func TestGetRepository(t *testing.T) {
	h := createHttpClient()
	repos := make([]types.Repository,0)
	h.ExecRequest("/api/repositories/", "GET", nil, &repos)
	fmt.Printf("%#v", repos)
	if len(repos) <=0 {
		t.Errorf("no repos")
	}
}

func TestCreateRepository(t *testing.T) {
	h := createHttpClient()
	repositoryDetails := types.CreateRepositoryDetails("local", "maven", "test-repo-cli", "simple-default")
	repositoryDetailsBytes, _ := json.Marshal(repositoryDetails)
	h.ExecRequest("/api/repositories/test-repo-cli", "PUT", repositoryDetailsBytes, nil)
}