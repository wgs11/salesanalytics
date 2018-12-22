package main

import (
	"fmt"
	"net/http"
)

type Review struct {
	day string `json: "date", db:"day"`
	store_id int `json: "store_id", db:"store_id"`
	outside []uint8 `json: "outside", db:"outside"`
	emp_sys []uint8 `json: "emp_sys", db:"emp_sys"`
	eating []uint8 `json: "eating", db:"eating"`
	merch []uint8 `json: "merch", db:"merch"`
	fountain []uint8 `json: "fountain", db:"fountain"`
	inventory  []uint8 `json: "inventory", db:"inventory"`
	backroom []uint8 `json: "backroom", db:"backroom"`
	restrooms []uint8 `json: "restrooms", db:"restrooms"`
	feedback string `json: "feedback", db:"feedback"`

}



func getReviews(w http.ResponseWriter, r *http.Request) {
	reviews, err := store.GetReviews("Wexford")
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(reviews[0].backroom)
}