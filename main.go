package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var slackClient Client
var data Data

type Data struct {
	ProjectData []ProjectData `json:"data"`
}

type ProjectData struct {
	Repository   string `json:"repository"`
	Message  string `json:"msg"`
	Channels []Channel `json:"slack_channels"`
}

func main() {
	slackClient = getClient("xoxb-834583945889-831079701299-54NDAFE9FyAI3NsmXF8AYjWe")
	data = getData("data.json")

	mux := http.NewServeMux()
	mux.HandleFunc("/gitlab-hook", mrHookHandler)

	log.Printf("listening on port %s", "8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func getData(fileName string) Data {
	var data Data

	jsonFile, err := os.Open(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &data)

	defer jsonFile.Close()

	return data
}