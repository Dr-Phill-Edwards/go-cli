package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetData(path string) []byte {
	uri := "https://api.github.com/" + path
	client := &http.Client{}
	request, _ := http.NewRequest("GET", uri, nil)
	request.Header.Add("Authorization", "token "+os.Getenv("GITHUB_ACCESS_TOKEN"))
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error " + err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body
}

func GetUser() map[string]interface{} {
	var user map[string]interface{}
	json.Unmarshal(GetData("user"), &user)
	return user
}

func GetRepos() []map[string]interface{} {
	var repos []map[string]interface{}
	json.Unmarshal(GetData("user/repos"), &repos)
	return repos
}
