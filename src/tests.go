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
