package main

import (
	"encoding/json"
	"net/http"
)

type RequestData struct {
	ObjectKind string `json:"object_kind"`
	Url string `json:"url"`
	WIP bool `json:"work_in_progress"`
	User struct {
		Name string `json:"name"`
	} `json:"user"`
	Project struct {
		Name string `json:"name"`
	}
	Repository struct {
		Name string `json:"name"`
		Homepage string `json:"homepage"`
	} `json:"repository"`
	Attributes struct {
		Action string  `json:"action"`
		Url string `json:"url"`
	} `json:"object_attributes"`
}

func mrHookHandler(w http.ResponseWriter, r *http.Request) {
	//TODO check that header contains Merge Request Hook
	if r.Method == "POST" {
		requestData := RequestData{}

		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		for _, project := range data.ProjectData {
			if project.Repository == requestData.Repository.Homepage &&
				requestData.Attributes.Action == "open" {
				for _, channel := range project.Channels {
					channel, _ := slackClient.getChannelByName(channel.Name)

					if err != nil {
						//TODO log error
					} else {
						slackClient.sendMessage(channel.ID, project.Message, requestData)
					}
				}

				break
			}
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

