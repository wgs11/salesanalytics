package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"html/template"
	"net/http"
	"strconv"
)
var fmap = template.FuncMap {
	"countStuff": countStuff,
	"stringify": stringify}

var templates = template.Must(template.New("").Funcs(fmap).ParseGlob("assets/*.gohtml"))
var cache = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))


func countStuff(numbers []uint8) int{
	total := 0
	for _, element := range numbers {
		total = total + int(element)
	}
	return total - (48 * len(numbers))
}

func stringify(number int) string {
	return strconv.Itoa(number)
}


func newRouter() *mux.Router {
	r := mux.NewRouter()
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/",http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler)
	r.HandleFunc("/", createPageHandler)
	r.HandleFunc("/signin", Signin)
	r.HandleFunc("/signup", Signup)
	r.HandleFunc("/days", getDaysHandler)
	r.HandleFunc("/day", getDayHandler).Methods("POST")
	r.HandleFunc("/employee", getEmployeeHandler)
	r.HandleFunc("/store/{location}", getPlaceHandler).Methods("GET")
	r.HandleFunc("/stores/{location}/reviews", getReviews)
	r.HandleFunc("/stores/{location}/reviews/{date}", getReview)
	r.HandleFunc("/admin", adminHandler)


	return r
}



func main() {

	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found.")
	} else {

		user := viper.GetString("production.user")
		dbname := viper.GetString("production.dbname")
		password := viper.GetString("production.password")
		connString := "user="+user+" dbname="+dbname+" password="+password+" sslmode=disable"
		fmt.Println(user,dbname,password)
		db, err := sql.Open("postgres", connString)
		if err != nil {
			panic(err)
		}
		err = db.Ping()
		if err != nil {
			panic(err)
		}
		InitStore(&dbStore{db: db})
		r := newRouter()
		http.ListenAndServe(":8080", r)
		fmt.Println("Serving on port 8080")
	}
}
