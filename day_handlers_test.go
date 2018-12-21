package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestGetDaysHandler(t *testing.T) {
	mockStore := InitMockStore()

	mockStore.On("GetDays").Return([]*Day{
		{"1990-01-02", 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10,10}}, nil).Once()
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(getDayHandler)
	hf.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := Day{"1990-01-02", 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10,10}
	d := []Day{}
	err = json.NewDecoder(recorder.Body).Decode(&d)
	if err != nil {
		t.Fatal(err)
	}
	actual := d[0]
	if actual != expected {
		t.Errorf("Handler returned unexpected body: got %v, got %v", actual, expected)

	}
	mockStore.AssertExpectations(t)

}

func TestCreateDaysHandler(t *testing.T) {
	mockStore := InitMockStore()

	mockStore.On("CreateDay", &Day{"1990-01-02", 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10,10}).Return(nil)
	form := newCreateDayForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(createDayHandler)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	mockStore.AssertExpectations(t)

}

func newCreateDayForm() *url.Values {
	form := url.Values{}
	form.Set("date", "1990-01-02")
	form.Set("ElevenAM", "10")
	form.Set("Noon", "10")
	form.Set("OnePM", "10")
	form.Set("TwoPM", "10")
	form.Set("ThreePM", "10")
	form.Set("FourPM", "10")
	form.Set("FivePM", "10")
	form.Set("SixPM", "10")
	form.Set("SevenPM", "10")
	form.Set("EightPM", "10")
	form.Set("NinePM", "10")
	form.Set("TenPM", "10")
	form.Set("ElevenPM", "10")
	form.Set("Total","10")
	return &form
}





