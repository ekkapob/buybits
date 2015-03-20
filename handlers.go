package main

import (
	"database/sql"
	"encoding/json"
	// "fmt"
	"github.com/ekkapob/buybits/model"
	"github.com/ekkapob/buybits/query"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

type httpfunc func(rw http.ResponseWriter, req *http.Request, db *sql.DB)

func dbHandler(fn httpfunc, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		fn(w, r, db)
	}
}

func writeJson(w http.ResponseWriter, data interface{}) {
	response, _ := json.Marshal(data)
	w.Write(response)
}

// func IndexHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	username := getSessionCookie(r)["username"]
// 	fmt.Println(username, len(username))

// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.Write([]byte("index"))
// }

func SignupHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	user := model.User{
		Username:  strings.TrimSpace(r.FormValue("username")),
		Password:  strings.TrimSpace(r.FormValue("password")),
		Firstname: strings.TrimSpace(r.FormValue("firstname")),
		Lastname:  strings.TrimSpace(r.FormValue("lastname")),
	}

	if len(user.Username) == 0 || len(user.Password) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		writeJson(w, model.Status{
			Error: "Username and password are required.",
		})
		return
	}

	if query.UserExist(db, user.Username) {
		writeJson(w, model.Status{
			Success: false,
			Error:   "Username is already taken.",
		})
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)
	err := query.AddUser(db, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJson(w, model.Status{
			Error: err.Error(),
		})
		return
	}

	writeJson(w, model.Status{
		Success: true,
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var (
		username = r.FormValue("username")
		password = r.FormValue("password")
		errmsg   = "Username or password is incorrect."
	)
	user, err := query.GetUser(db, username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJson(w, model.Status{
			Error: errmsg,
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		writeJson(w, model.Status{
			Error: errmsg,
		})
		return
	}
	setCookie(w, user.Username)
	writeJson(w, model.LoginResponse{
		Status: model.Status{Success: true},
		User:   user,
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	clearCookie(w)
	writeJson(w, model.Status{
		Success: true,
	})
}

func PostsHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	posts, err := query.GetPosts(db, 30, 0)
	if err != nil {
		writeJson(w, model.Status{})
		return
	}
	writeJson(w, model.PostsResponse{
		Status: model.Status{Success: true},
		Posts:  posts,
	})
}
