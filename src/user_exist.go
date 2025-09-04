package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func userExistHandler(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/data/userexist?user=nate&cookie=f
	var cookie = r.URL.Query().Get("cookie")
	var name = r.URL.Query().Get("user")
	if !NameCookieAuth(name, cookie) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("Failed Cookie Auth"))
		return
	}
	files, err := os.ReadDir("../data/saves")
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		fmt.Println("error in readdir:", err)
		return
	}
	var doesExist bool = false
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		f := strings.TrimSuffix(file.Name(), ".json")
		if f == name {
			doesExist = true
		}
	}
	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	if doesExist {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}
