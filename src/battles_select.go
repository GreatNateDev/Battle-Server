package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var payload = make(map[string]map[string]interface{})

func battleSelectHandler(w http.ResponseWriter, r *http.Request) {
	// Example Payload http://localhost:8080/data/battles?user=nate,cookie=fbgdhsjeui
	var cookie = r.URL.Query().Get("cookie")
	var name = r.URL.Query().Get("user")
	if !NameCookieAuth(name, cookie) {
		fmt.Println("AUTH FAILED")
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("Failed Cookie Auth"))
		return
	}
	files, err := os.ReadDir("../data/battles")
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		fmt.Println("error in readdir:", err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Read the JSON file
		path := filepath.Join("../data/battles", file.Name())
		f, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("file read err:", err)
			continue
		}
		var battle map[string]interface{}
		if err := json.Unmarshal(f, &battle); err != nil {
			fmt.Println("unmarshal err:", err)
			continue
		}
		key := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		payload[key] = battle
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(payload)
}
