package artapi

import (
	"net/http"
	"fmt"
	"github.com/spf13/viper"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"strconv"
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
	h.url = viper.GetString("url")
	h.user = viper.GetString("user")
	h.password = viper.GetString("password")
	return h
}

func (h *HttpClient) ExecRequest(path string, method string, payload []byte, queryRes interface{}) {

	if viper.GetBool("verbose") {
		fmt.Printf("ARTIFACTORY_URL: [%s]\n", h.url)
		fmt.Printf("ARTIFACTORY_USER: [%s]\n", h.user)
		fmt.Printf("ARTIFACTORY_PASSWORD: [%s]\n", "***************")
	}

	if viper.GetString("verbose") == "true" {
		fmt.Printf("url: [%s], path [%s], payload [%s]", h.url, path, string(payload))
	}
	req, err := http.NewRequest(method, h.url+path, bytes.NewReader(payload))
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(h.user, h.password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// read the response body to a variable
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	//print raw response body for debugging purposes
	if viper.GetBool("verbose") {
		fmt.Println("\n\n", bodyString, "\n\n")
	}

	if resp.StatusCode != 200 {
		panic("response status code is different than 200: " + strconv.Itoa(resp.StatusCode))
	}

	//reset the response body to the original unread state
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	if queryRes != nil {
		err = json.NewDecoder(resp.Body).Decode(queryRes)
		if err != nil {
			panic(err)
		}
	}

}
