package main

import (
	"fmt"
	"log"
	"net/http"
)

type Review struct {
	day string `json: "date", db:"day"`
	store_id int `json: "store_id", db:"store_id"`
	Outside []uint8 `json: "outside", db:"outside"`
	Emp_sys []uint8 `json: "emp_sys", db:"emp_sys"`
	Eating []uint8 `json: "eating", db:"eating"`
	Merch []uint8 `json: "merch", db:"merch"`
	Fountain []uint8 `json: "fountain", db:"fountain"`
	Inventory  []uint8 `json: "inventory", db:"inventory"`
	Backroom []uint8 `json: "backroom", db:"backroom"`
	Restrooms []uint8 `json: "restrooms", db:"restrooms"`
	Feedback string `json: "feedback", db:"feedback"`

}



func getReviews(w http.ResponseWriter, r *http.Request) {
	reviews, err := store.GetReviews("Wexford")
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	reviewblock := []Review{}
	for _,element := range reviews {
		reviewblock = append(reviewblock, (*element))
	}
	fmt.Println(reviewblock[0])
	err = templates.ExecuteTemplate(w, "review", reviewblock)
	if err != nil {
		log.Fatal(err)
	}
}