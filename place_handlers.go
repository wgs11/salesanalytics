package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Place struct {
	ID int `json: "id"`
	Location string `json: "name"`
	SMID int `json: "smid"`
}



func getPlaceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	location, _ := store.GetPlace(vars["location"])
	fmt.Println(location)

}
