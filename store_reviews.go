package main

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
		if err := rows.Scan(&review.day, &ans, &review.Feedback); err != nil {
			return nil, err
		}
		review.Outside = ans[:7]
		review.Emp_sys = ans[7:14]
		review.Eating = ans[14:30]
		review.Merch = ans[30:46]
		review.Fountain = ans[46:60]
		review.Inventory = ans[60:74]
		review.Backroom = ans[74:88]
		review.Restrooms = ans[88:100]
		reviews = append(reviews, review)
		reviews = append(reviews, review)
		reviews = append(reviews, review)
	}
	return reviews, nil
}