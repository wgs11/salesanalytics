package main

import (
	"strconv"
)

//REVIEW STRUCT//
//type Review struct {
//	date string `json: "date", db:"day"`
//	store_id int `json: "sid", db:"store_id"`
//	answers bitarray.BitArray `json: "answers", db:"answers"`
//	feedback string `json: "feedback", db:"feedback"`
//}

func (store *dbStore) GetReview(location string, date string) (*Review, error) {
	review := &Review{}
	row, err := store.db.Query("SELECT day, answers::bit(100), feedback FROM reviews WHERE day = $1 and store_id = $2", date, location)
	if err != nil {
		return nil, err
	} else {
		defer row.Close()
		err := row.Next()
		if err {
			ans := []uint8{}
			if err := row.Scan(&review.Day, &ans, &review.Feedback); err != nil {
				return nil, err
			}
			review.Day = review.Day[:10]
			review.Outside = ans[:7]
			review.Emp_sys = ans[7:14]
			review.Eating = ans[14:30]
			review.Merch = ans[30:46]
			review.Fountain = ans[46:60]
			review.Inventory = ans[60:74]
			review.Backroom = ans[74:88]
			review.Restrooms = ans[88:100]
			return review, nil
		}
	}

	return nil, nil
}

func (store *dbStore) GetReviews(location string) ([]*Review, error) {
	rows, err := store.db.Query("SELECT day, answers::bit(100), feedback FROM reviews WHERE store_id = $1", location)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reviews := []*Review{}
	for rows.Next() {
		review := &Review{}
		review.Store_id, _ = strconv.Atoi(location)
		ans := []uint8{}
		if err := rows.Scan(&review.Day, &ans, &review.Feedback); err != nil {
			return nil, err
		}
		review.Day = review.Day[:10]
		review.Outside = ans[:7]
		review.Emp_sys = ans[7:14]
		review.Eating = ans[14:30]
		review.Merch = ans[30:46]
		review.Fountain = ans[46:60]
		review.Inventory = ans[60:74]
		review.Backroom = ans[74:88]
		review.Restrooms = ans[88:100]
		reviews = append(reviews, review)
	}

	return reviews, nil
}