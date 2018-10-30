package controllers

import (
	"encoding/json"
	u "cashgone/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type Github struct {
	TotalCount int `json:"total_count"`
	IncompleteResults bool `json:"incomplete_results"`
	Items []Item `json:"items"`
}

type Item struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

var SearchRepos = func(w http.ResponseWriter, r *http.Request) {
	github := &Github{}
	params := mux.Vars(r)
	request, _ := http.NewRequest("GET", "https://api.github.com/search/repositories?q=" + params["query"], nil)
	request.Header.Add("accept", "application/json")
	request.Header.Add("content-type", "application/json")
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()

	err := json.NewDecoder(response.Body).Decode(github)
	if err != nil {
		u.Respond(w, http.StatusForbidden, u.Message(false, "Error while loading github api."))
		return
	}
	resp := u.Message(true, "success")
	resp["data"] = github
	u.Respond(w, http.StatusOK, resp)
}