package main

import "fmt"

//REVIEW STRUCT//
//type Review struct {
//	date string `json: "date", db:"day"`
//	store_id int `json: "sid", db:"store_id"`
//	answers bitarray.BitArray `json: "answers", db:"answers"`
//	feedback string `json: "feedback", db:"feedback"`
//}
func (store *dbStore) GetReviews(location string) ([]*Review, error) {
	rows, err := store.db.Query("SELECT day, answers, feedback FROM reviews WHERE store_id = 1")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reviews := []*Review{}
	for rows.Next() {
		review := &Review{}
		ans := []uint8{}
		if err := rows.Scan(&review.day, &ans, &review.feedback); err != nil {
			return nil, err
		}
		review.outside = ans[:7]
		review.emp_sys = ans[7:14]
		review.eating = ans[14:30]
		review.merch = ans[30:46]
		review.fountain = ans[46:60]
		review.inventory = ans[60:74]
		review.backroom = ans[74:88]
		review.restrooms = ans[88:100]
		fmt.Println(review.outside)
		fmt.Println(review.emp_sys)
		fmt.Println(ans)
		reviews = append(reviews, review)
	}
	return reviews, nil
}