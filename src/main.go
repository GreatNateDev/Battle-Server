package main

import (
	"fmt"
	"net/http"
)

const IP = "127.0.0.1"
const PORT = "8080"

var COOKIES []string
var TMP_COOKIES []string

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Main Page Hit")
}

func apiSetup() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/savedata", saveHandler)
	http.HandleFunc("/submitcookie", cookieHandler)
	http.HandleFunc("/data/pokemon", pokemonDataHandler)
}

func main() {
	apiSetup()
	cookieSetter()
	http.ListenAndServe(IP+":"+PORT, nil)
}
