package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
}

func getUser(w http.ResponseWriter, r *http.Request) string {
	session, _ := cache.Get(r, "cookie-name")
	if val, ok := session.Values["user"].(string); ok {
		return val
	} else {
		return ""
	}
}

func IsSignedIn(w http.ResponseWriter, r *http.Request) bool {
	session, _ := cache.Get(r, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return false
	} else {return true}
}

func Signup(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(creds.Username, creds.Password)
	store.CreateUser(creds)
}

func isAdmin(w http.ResponseWriter, r *http.Request) bool {
	session, _ := cache.Get(r, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		fmt.Println(session.Values["authenticated"])
		fmt.Println("signed in and authenticated")
		if admin, k := session.Values["admin"].(bool); k && admin {
			return true
		} else {fmt.Println(admin, k)}
	}
	return false
}

func Signin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	creds := &Credentials{}
	fmt.Println(r.FormValue("username"))
	fmt.Println(r.FormValue("password"))
	creds.Username = r.FormValue("username")
	creds.Password = r.FormValue("password")
	if store.CheckUser(creds) == nil {
		session,_ := cache.Get(r, "cookie-name")
		session.Values["authenticated"] = true
		session.Values["user"] = creds.Username
		fmt.Println(creds.Username)
		if (creds.Username == "Sheppy") {
			fmt.Println("user is admin")
			session.Values["admin"] = true
		} else { session.Values["admin"] = false}
		session.Save(r,w)
		fmt.Println(session.Values["admin"])
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		fmt.Println("/")
	}
}
