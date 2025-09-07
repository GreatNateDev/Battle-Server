package main

import (
	"fmt"
	"net/http"
)

const (
	DEBUG = true
	IP    = "127.0.0.1"
	PORT  = "8080"
)

var (
	COOKIES     []string
	TMP_COOKIES []string
	CookieMap   = make(map[string]string)
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Main Page Hit")
}

func apiSetup() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/savedata", saveHandler)
	http.HandleFunc("/submitcookie", cookieHandler)
	http.HandleFunc("/test/pokemon", pokemonDataHandler)
	http.HandleFunc("/test/auth", nameCookieHandler)
	http.HandleFunc("/data/spritedl", spriteDownloadHandler)
	http.HandleFunc("/data/battles", battleSelectHandler)
	http.HandleFunc("/data/userexist", userExistHandler)
}

func main() {
	apiSetup()
	cookieSetter()
	fmt.Println("Serving Poke_Backend on port :8080")
	http.ListenAndServe(IP+":"+PORT, nil)
}
