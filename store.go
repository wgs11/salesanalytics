package main

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Store interface {
	CreateDay(day *Day) error
	GetDays() ([]*Day, error)
	GetDay(estimate float32) ([]*Day, error)
	GetEmployee(id int) ([]*Employee, error)
	GetPlace(location string) (*Place, error)
	CreateUser(creds *Credentials) error
	CheckUser(creds *Credentials) error
	GetReviews(location string) ([]*Review, error)
}


type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CheckUser(creds *Credentials) error {
	dummyCreds := &Credentials{}
	row, err := store.db.Query("SELECT password FROM users WHERE username = $1", creds.Username)
	if err != nil {
		return err } else {
		defer row.Close()
		err := row.Next()
		if err {
			row.Scan(&dummyCreds.Password)
			fmt.Println(dummyCreds.Password)
			return(bcrypt.CompareHashAndPassword([]byte(dummyCreds.Password), []byte(creds.Password)))
		}
	}
	return nil
}
func (store *dbStore) CreateUser(creds *Credentials) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.Password),8)
	fmt.Println(creds.Username, creds.Password)
	_, err := store.db.Query("INSERT INTO users VALUES ($1, $2)", string(creds.Username), string(hashedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (store *dbStore) GetPlace(location string) (*Place, error) {
	place := &Place{}
	row, err := store.db.Query("SELECT id FROM stores WHERE location = $1", location)
	if err != nil {
		return nil, err
	} else {
		defer row.Close()
		err := row.Next()
		if err {
			place.Location = location
			row.Scan(&place.ID)
		}
	}

	return place, nil
}

func (store *dbStore) GetEmployee(id int) ([]*Employee, error) {
	row, err := store.db.Query("SELECT FN,LN,Role FROM employees WHERE ID = $1", id)
	if err != nil {
		return nil, err
	}
	employees := []*Employee{}
	defer row.Close()
	employee := &Employee{}
	employee.ID = id
	for row.Next() {
		var temp int
		if err := row.Scan(&employee.FN, &employee.LN, &temp); err != nil {
			fmt.Println(err)
			return nil, err
		}
		switch (temp) {
		case 4: employee.Role = "Store Manager"
		case 3: employee.Role = "Assistant Manager"
		case 2: employee.Role = "Team Manager"
		case 1: employee.Role = "Team Leader"
		case 0: employee.Role = "Team Member"
		}
	}
	employees = append(employees,employee)
	return employees, nil
}


func (store *dbStore) GetDay(estimate float32) ([]*Day, error) {
	var estimate_range float32 = 400.00
	var half = estimate_range / 2.00
	row, err := store.db.Query("SELECT ROUND(AVG(ElevenAM),2) ElevenAM, ROUND(AVG(Noon),2) Noon, ROUND(AVG(OnePM),2) OnePM, ROUND(AVG(TwoPM),2) TwoPM, ROUND(AVG(ThreePM),2) ThreePM, ROUND(AVG(FourPM),2) FourPM, ROUND(AVG(FivePM),2) FivePM, ROUND(AVG(SixPM),2) SixPM, ROUND(AVG(SevenPM),2) SevenPM, ROUND(AVG(EightPM),2) EightPM, ROUND(AVG(NinePM),2) NinePM, ROUND(AVG(TenPM),2) TenPM, ROUND(AVG(ElevenPM),2) ElevenPM, ROUND(AVG(Total),2) Total from days where Total Between $1 and $2",(estimate-half), (estimate+half))
	if err != nil {
		return nil, err
	}
	days := []*Day{}
	defer row.Close()
	fmt.Println(row)
	day := &Day{}
	day.Date = "1990-01-01"
	for row.Next() {
		if err := row.Scan(&day.ElevenAM, &day.Noon, &day.OnePM, &day.TwoPM, &day.ThreePM, &day.FourPM, &day.FivePM, &day.SixPM, &day.SevenPM, &day.EightPM, &day.NinePM, &day.TenPM, &day.ElevenPM, &day.Total); err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	days = append(days, day)
	day = &Day{}
	day.Date = "1990-01-01"
	row, err = store.db.Query("SELECT MAX(ElevenAM) ElevenAM, MAX(Noon) Noon, MAX(OnePM) OnePM, MAX(TwoPM) TwoPM, MAX(ThreePM) ThreePM, MAX(FourPM) FourPM, MAX(FivePM) FivePM, MAX(SixPM) SixPM, MAX(SevenPM) SevenPM, MAX(EightPM), MAX(NinePM), MAX(TenPM), MAX(ElevenPM), MAX(Total) FROM days WHERE Total BETWEEN $1 AND $2",(estimate-half),(estimate+half))
	if err != nil {
		return nil, err
	}
	for row.Next() {
		if err := row.Scan(&day.ElevenAM, &day.Noon, &day.OnePM, &day.TwoPM, &day.ThreePM, &day.FourPM, &day.FivePM, &day.SixPM, &day.SevenPM, &day.EightPM, &day.NinePM, &day.TenPM, &day.ElevenPM, &day.Total); err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	days = append(days,day)
	day = &Day{}
	day.Date = "1990-01-01"
	row, err = store.db.Query("SELECT MIN(ElevenAM) ElevenAM, MIN(Noon) Noon, MIN(OnePM) OnePM, MIN(TwoPM) TwoPM, MIN(ThreePM) ThreePM, MIN(FourPM) FourPM, MIN(FivePM) FivePM, MIN(SixPM) SixPM, MIN(SevenPM) SevenPM, MIN(EightPM), MIN(NinePM), MIN(TenPM), MIN(ElevenPM), MIN(Total) FROM days WHERE Total BETWEEN $1 AND $2",(estimate-half),(estimate+half))
	if err != nil {
		return nil, err
	}
	for row.Next() {
		if err := row.Scan(&day.ElevenAM, &day.Noon, &day.OnePM, &day.TwoPM, &day.ThreePM, &day.FourPM, &day.FivePM, &day.SixPM, &day.SevenPM, &day.EightPM, &day.NinePM, &day.TenPM, &day.ElevenPM, &day.Total); err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	days = append(days,day)
	return days, nil
}


func (store *dbStore) CreateDay(day *Day) error {
	fmt.Println(day.Date, day.ElevenAM, day.Noon, day.OnePM, day.TwoPM, day.ThreePM, day.FourPM, day.FivePM, day.SixPM, day.SevenPM, day.EightPM, day.NinePM, day.TenPM, day.ElevenPM, day.Total)
	_, err := store.db.Query("INSERT INTO days(Date, ElevenAM, Noon, OnePM, TwoPM, ThreePM, FourPM, FivePM, SixPM, SevenPM, EightPM, NinePM, TenPM, ElevenPM, Total) VALUES ($1::date, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)",
		day.Date, day.ElevenAM, day.Noon, day.OnePM, day.TwoPM, day.ThreePM, day.FourPM, day.FivePM, day.SixPM, day.SevenPM, day.EightPM, day.NinePM, day.TenPM, day.ElevenPM, day.Total)
	return err
}


func (store *dbStore) GetDays() ([]*Day, error) {
	rows, err := store.db.Query("SELECT date_trunc('hour',Date), ElevenAM, Noon, OnePM, TwoPM, ThreePM, FourPM, FivePM, SixPM, SevenPM, EightPM, NinePM, TenPM, ElevenPM, Total from days")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	days := []*Day{}
	for rows.Next() {
		day := &Day{}

		if err := rows.Scan(&day.Date, &day.ElevenAM, &day.Noon, &day.OnePM, &day.TwoPM, &day.ThreePM, &day.FourPM, &day.FivePM, &day.SixPM, &day.SevenPM, &day.EightPM, &day.NinePM, &day.TenPM, &day.ElevenPM, &day.Total); err != nil {
			return nil, err
		}
		days = append(days, day)
	}
	return days, nil
}


var store Store


func InitStore(s Store) {
	store = s
}
