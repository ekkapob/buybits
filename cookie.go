package main

import (
	"github.com/gorilla/securecookie"
	"net/http"
)

var (
	// Generated by securecookie.GenerateRandomKey(64) and (32)
	hashKey  = []byte{144, 70, 174, 35, 6, 213, 43, 118, 83, 230, 186, 64, 107, 144, 8, 112, 176, 2, 0, 88, 81, 226, 52, 100, 242, 56, 45, 57, 221, 60, 129, 68, 63, 214, 63, 85, 59, 229, 119, 144, 48, 246, 218, 149, 198, 172, 95, 50, 1, 152, 94, 238, 9, 114, 62, 55, 87, 35, 54, 164, 119, 125, 245, 85}
	blockKey = []byte{199, 98, 23, 101, 194, 140, 193, 192, 187, 97, 59, 164, 166, 40, 165, 195, 14, 153, 188, 255, 24, 255, 222, 7, 36, 103, 55, 174, 182, 107, 174, 153}
	scookie  = securecookie.New(hashKey, blockKey)
)

func setCookie(w http.ResponseWriter, username string) {
	cookieValue := map[string]string{
		"username": username,
	}
	encoded, err := scookie.Encode("session", cookieValue)
	if err != nil {
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: encoded,
		Path:  "/",
	})
}

func clearCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}

func getSessionCookie(r *http.Request) (sessioncookie map[string]string) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return sessioncookie
	}
	scookie.Decode("session", cookie.Value, &sessioncookie)
	return sessioncookie
}
