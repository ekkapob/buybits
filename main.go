package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db := getDb()
	defer db.Close()

	r := mux.NewRouter()
	// r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/signup", dbHandler(SignupHandler, db)).Methods("POST")
	r.HandleFunc("/login", dbHandler(LoginHandler, db)).Methods("POST")
	r.HandleFunc("/logout", LogoutHandler).Methods("POST")

	r.HandleFunc("/posts", dbHandler(PostsHandler, db))
	http.ListenAndServe(":8080", r)
}
