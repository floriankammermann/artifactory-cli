package artapi

import (
	"net/http"
	"fmt"
	"github.com/spf13/viper"
	"encoding/json"
)

var h *HttpClient

type HttpClient struct {
	url string
	user string
	password string
	verbose bool
}

func InitHttpClient() *HttpClient {
	h = new(HttpClient)
	return h
}

func (h *HttpClient) ExecRequest(path string, queryRes interface{}) {

	if viper.GetString("verbose") == "true" {
		fmt.Printf("url: [%s], path [%s]", h.url, path)
	}

	req, err := http.NewRequest("GET", h.url+path, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(queryRes)
	if err != nil {
		panic(err)
	}
}
