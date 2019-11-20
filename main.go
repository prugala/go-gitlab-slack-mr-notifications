package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
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
	Repository string    `json:"repository"`
	Message    string    `json:"msg"`
	Channels   []Channel `json:"slack_channels"`
}

func init() {
	godotenv.Load(".env")
}

func main() {
	slackClient = getClient(os.Getenv("SLACK_TOKEN"))
	data = getData("data.json")

	mux := http.NewServeMux()
	mux.HandleFunc("/gitlab-hook", mrHookHandler)

	log.Printf("listening on port %s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), mux))
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
