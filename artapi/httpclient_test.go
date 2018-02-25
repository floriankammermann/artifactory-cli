package artapi

import (
	"testing"
	"os"
	"github.com/floriankammermann/artifactory-cli/types"
	"fmt"
)


func TestGetRepository(t *testing.T) {

	h := InitHttpClient()
	h.user = os.Getenv("USERNAME")
	h.password = os.Getenv("USERNAME")
	h.url = os.Getenv("URL")
	repos := make([]types.Repository,0)
	h.ExecRequest("/api/repositories/", &repos)
	fmt.Printf("%#v", repos)
	if len(repos) <=0 {
		t.Errorf("no repos")
	}

}