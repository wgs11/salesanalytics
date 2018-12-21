package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
)

type Day struct {
	Date		string `json: "date"`
	ElevenAM	float32 `json: "ElevenAM"`
	Noon		float32 `json: "noon"`
	OnePM		float32 `json: "onePM"`
	TwoPM		float32 `json: "twoPM"`
	ThreePM 	float32 `json: "threePM"`
	FourPM 		float32 `json: "fourPM"`
	FivePM 		float32 `json: "fivePM"`
	SixPM 		float32 `json: "sixPM"`
	SevenPM 	float32 `json: "sevenPM"`
	EightPM 	float32 `json: "eightPM"`
	NinePM 		float32 `json: "ninePM"`
	TenPM 		float32 `json: "tenPM"`
	ElevenPM 	float32 `json: "elevenPM"`
	Total		float32 `json: "total"`

}

type DaysStruct struct {
	PageTitle string
	DayAvg Day
	DayMax Day
	DayMin Day
}



func stringTo32Float(sales string) float32 {
	fval, err := strconv.ParseFloat(sales,32)
	if err != nil {
		fmt.Println(err)
	}
	return float32(fval)
}
//1st line is already properly date formatted, so we're just making a Day to return to the Handler
func parseFile(file multipart.File) Day {

	scanner := bufio.NewScanner(file)
	curline := 0
	day := Day{}
	for scanner.Scan() {
		if curline == 0 {
			day.Date = scanner.Text()
		} else {
			value := stringTo32Float(scanner.Text())
			switch curline {
			case 1: day.ElevenAM = value
			case 2: day.Noon = value
			case 3: day.OnePM = value
			case 4: day.TwoPM = value
			case 5: day.ThreePM = value
			case 6: day.FourPM = value
			case 7: day.FivePM = value
			case 8: day.SixPM = value
			case 9: day.SevenPM = value
			case 10: day.EightPM = value
			case 11: day.NinePM = value
			case 12: day.TenPM = value
			case 13: day.ElevenPM = value
			case 14: day.Total = value
			}
		}

		fmt.Println(scanner.Text())
		curline += 1
	}
	return day
}

func getDayHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	start := r.FormValue("start_date")
	end := r.FormValue("end_date")
	estimate := r.FormValue("estimate")
	fmt.Println(start)
	fmt.Println(end)
	fmt.Println(estimate)
	days, err := store.GetDay(stringTo32Float(estimate))
	Days := DaysStruct{}
	Days.DayAvg = (*days[0])
	Days.DayMax = (*days[1])
	Days.DayMin = (*days[2])
	Days.PageTitle = "This is a Title"
	err = templates.ExecuteTemplate(w,"layout",Days)
	if err != nil {
		log.Fatal("Cannot retrive page.", err)
	}
	if err != nil {
		fmt.Println(fmt.Errorf("Error %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}


func getDaysHandler (w http.ResponseWriter, r *http.Request) {
	days, err := store.GetDays()
	dayListBytes, err := json.Marshal(days)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(dayListBytes)
}


func createDayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(100000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		m := r.MultipartForm
		files := m.File["q"]
		day := Day{}
		for i, _ := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			day = parseFile(file)
			fmt.Println(&day)
			err = store.CreateDay(&day)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func createPageHandler(w http.ResponseWriter, r *http.Request) {
	if !IsSignedIn(w,r){
		err := templates.ExecuteTemplate(w, "login", "")
		if err != nil {
			log.Fatal("Cannot retrieve login page.")
		}
	} else {
		err := templates.ExecuteTemplate(w, "home", "")
		if err != nil {
			log.Fatal("Cannot retrive page.", err)
		}
	}

}