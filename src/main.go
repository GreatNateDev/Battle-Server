package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const IP = "127.0.0.1"
const PORT = "8080"

var COOKIES []string
var TMP_COOKIES []string

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Main Page Hit")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("pong")
	w.Write([]byte("pong"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login Page Hit")
	user := r.URL.Query().Get("user")
	pass := r.URL.Query().Get("pass")
	f, err := os.ReadFile("../data/logins.json")
	if err != nil {
		panic(err)
	}
	var logins map[string]interface{}
	json.Unmarshal(f, &logins)
	stored, ok := logins[user] // stored is the value for user ("password1"), ok is true if found
	if !ok {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}
	if pass != stored.(string) { // type assert to string since interface{} is generic
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	fmt.Println("Login successful!")
	cookie := randomString(10)
	w.Write([]byte(cookie))
	fmt.Println("Sent cookie to authenticated client: " + cookie)
	COOKIES = append(COOKIES, cookie)
	TMP_COOKIES = append(TMP_COOKIES, cookie)
}

func cookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := r.URL.Query().Get("cookie")
	fmt.Println("Received Cookie: " + cookie)
	if !contains(TMP_COOKIES, cookie) {
		TMP_COOKIES = append(TMP_COOKIES, cookie)
	}
}

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

func pokemonDataHandler(w http.ResponseWriter, r *http.Request) {
	mon := r.URL.Query().Get("mon")

	fmt.Println("Pokemon Data Requested { " + mon + " }")
	filePath := fmt.Sprintf("../data/pokemon/%s/data.json", mon)
	http.ServeFile(w, r, filePath)
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

func cookieSetter() {
	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for range ticker.C {
			cookieKill()
		}
	}()
}

func cookieKill() {
	fmt.Println("Clearing Cookies")
	fmt.Printf("Current Cookies: %v TMP_COOKIES: %v\n", COOKIES, TMP_COOKIES)
	// Build a new list to keep cookies that are also in TMP_COOKIES:
	var keptCookies []string
	for _, cookie := range TMP_COOKIES {
		if contains(COOKIES, cookie) {
			keptCookies = append(keptCookies, cookie)
		}
		// else: do nothing – we're "removing" by not appending
	}

	COOKIES = keptCookies
	TMP_COOKIES = nil // clear TMP_COOKIES
}

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// TODO pair cookies to username on dish
