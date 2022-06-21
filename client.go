package main

import (
	"fmt"
	"encoding/json"
	"path/filepath"
	"gdrive/auth"
)

const clientId = "801306755384-43n67frda30ohcenp9mmokfj6ijettvm.apps.googleusercontent.com"
const clientSecret = "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"

type _client struct {
	ClientId string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}


func GetClient() _client {
        clientPath := getDefaultConfigFile()
	content, exists, err := auth.ReadFile(clientPath)

	client := _client{
		ClientId: clientId,
  		ClientSecret: clientSecret,
	}

	if err != nil || exists == false {
		return client
	}

	err2 := json.Unmarshal(content, &client)
	if err2 != nil {
		fmt.Println("client.json parse error.")
	}
	return client
}

// util.go and mod
func getDefaultConfigFile() string {
	return filepath.Join(Homedir(), ".gdrive", "client.json")
}

