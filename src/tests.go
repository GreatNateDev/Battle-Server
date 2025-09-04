package main

import (
	"fmt"
	"net/http"
)

func pokemonDataHandler(w http.ResponseWriter, r *http.Request) {
	mon := r.URL.Query().Get("mon")

	fmt.Println("Pokemon Data Requested { " + mon + " }")
	filePath := fmt.Sprintf("../data/pokemon/%s/data.json", mon)
	http.ServeFile(w, r, filePath)
}
func nameCookieHandler(w http.ResponseWriter, r *http.Request) {
	var cookie = r.URL.Query().Get("cookie")
	var name = r.URL.Query().Get("user")
	if !NameCookieAuth(name, cookie) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("Failed Cookie Auth"))
		return
	}
}
