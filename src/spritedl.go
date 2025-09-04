package main

import (
	"fmt"
	"net/http"
	"os"
)

func spriteDownloadHandler(w http.ResponseWriter, r *http.Request) {
	// Example Request http://localhost:8080/data/spritedl?user=nate&cookie=abcdefghij&sprite=pikachu&side=front
	var cookie = r.URL.Query().Get("cookie")
	var name = r.URL.Query().Get("user")
	var sprite = r.URL.Query().Get("sprite")
	var side = r.URL.Query().Get("side")
	if !NameCookieAuth(name, cookie) {
		fmt.Println("AUTH FAILED")
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("Failed Cookie Auth"))
		return
	}
	f, err := os.ReadFile("../data/pokemon/" + sprite + "/" + side + ".png")
	if err != nil {
		fmt.Println("spritenotfound")
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write([]byte(f))
}
