package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func saveHandler(w http.ResponseWriter, r *http.Request) {
	// Example Usage:
	// http://server.local:8080/savedata?name=root&type=store&cookie=XXX&data={json}
	// http://server.local:8080/savedata?name=root&type=download&cookie=XXX

	playerName := r.URL.Query().Get("name")
	typeof := r.URL.Query().Get("type")
	data := r.URL.Query().Get("data")
	cookie := r.URL.Query().Get("cookie")

	// cookie auth
	if !NameCookieAuth(playerName, cookie) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("Failed Cookie Auth"))
		return
	}

	// try reading an existing save (not required for store, but kept for consistency)
	f, err := os.ReadFile("../data/saves/" + playerName + ".json")
	if err != nil {
		fmt.Println("No existing save or read error:", err)
	}

	var file map[string]interface{}
	_ = json.Unmarshal(f, &file) // safe ignore if empty

	switch typeof {
	case "store":
		// remove old save
		os.Remove("../data/saves/" + playerName + ".json")

		// parse incoming JSON string
		var parsed interface{}
		if err := json.Unmarshal([]byte(data), &parsed); err != nil {
			http.Error(w, "Invalid JSON in data param", http.StatusBadRequest)
			return
		}

		// re-marshal for pretty formatting
		jsonBytes, err := json.MarshalIndent(parsed, "", "  ")
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		// write new save
		if err := os.WriteFile("../data/saves/"+playerName+".json", jsonBytes, 0644); err != nil {
			http.Error(w, "Error writing file", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Save stored successfully"))

	case "download":
		f, err = os.ReadFile("../data/saves/" + playerName + ".json")
		if err != nil {
			http.Error(w, "Error: No Save", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(f)

	default:
		http.Error(w, "Invalid type parameter", http.StatusBadRequest)
	}
}
