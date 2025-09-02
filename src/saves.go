package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func saveHandler(w http.ResponseWriter, r *http.Request) {
	//Example Useage
	//http://server.local:8080/savedata?name=root&type=store&[cookie]&data={json} this is for uploading a save
	//http://server.local:8080/savedata?name=root&type=download&[cookie]&data={json} <-- (this part doesnt matter it will be discarded) this is for downloading a save
	playerName := r.URL.Query().Get("name")
	typeof := r.URL.Query().Get("type")
	data := r.URL.Query().Get("data")
	//cookie := r.URL.Query().Get("cookie")
	f, err := os.ReadFile("../data/saves/" + playerName + ".json")
	if err != nil {
		fmt.Println("err")
	}
	var file map[string]interface{}
	json.Unmarshal(f, &file)
	switch typeof {
	case "store":
		os.Remove("../data/saves/" + playerName + ".json")
		jsonBytes, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Println("error")
		}
		err = os.WriteFile("../data/saves/"+playerName+".json", jsonBytes, 0644)
		if err != nil {
			fmt.Println("error")
		}
	case "download":
		f, err = os.ReadFile("../data/saves/" + playerName + ".json")
		if err == os.ErrNotExist {
			w.Write([]byte("Error: No Save"))
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(f))
		}
	}

}
