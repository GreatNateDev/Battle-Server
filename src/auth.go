package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

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
	CookieMap[user] = cookie
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
	set := make(map[string]struct{})
	for _, v := range COOKIES {
		set[v] = struct{}{}
	}

	// Filter the dict
	for k, v := range CookieMap {
		if _, exists := set[v]; !exists {
			delete(CookieMap, k)
		}
	}
}
func NameCookieAuth(name string, cookie string) bool {
	fmt.Println("Starting Auth:", name, cookie)
	if DEBUG == true {
		if cookie == "abcdefghij" {
			return true
		}
	}
	stored, ok := CookieMap[name]
	if !ok {
		return false
	}
	return cookie == stored
}
