package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)


type StoreSuite struct {
	suite.Suite
	store *dbStore
	db *sql.DB
}

func (s *StoreSuite) SetupSuite() {
	connString :="user=postgres dbname=sales_analytics password=wrong1 sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	s.store = &dbStore{db: db}
}

func (s *StoreSuite) SetupTest() {

	_, err := s.db.Query("DELETE FROM days")
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *StoreSuite) TearDownSuite() {
	s.db.Close()
}

func TestStoreSuite(t *testing.T) {
	s := new(StoreSuite)
	suite.Run(t,s)
}

func (s *StoreSuite) TestCreateDay() {
	s.store.CreateDay(&Day{
		Date: "1990-01-02",
		ElevenAM: 20,
		Noon: 200,
		OnePM: 100,
		TwoPM: 200,
		ThreePM: 140,
		FourPM: 200,
		FivePM: 0,
		SixPM: 500,
		SevenPM: 223,
		EightPM: 90,
		NinePM: 432,
		TenPM: 78,
		ElevenPM: 0,
		Total: 900,})
	res, err := s.db.Query(`SELECT COUNT(*) FROM days WHERE Date='1990-01-02'`)
	if err != nil {
		s.T().Fatal(err)
	}
	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Error(err)
		}
	}

	if count != 1 {
		s.T().Errorf("Incorrect count, wanted 1, got %d", count)
	}
}



func (s *StoreSuite) TestGetDay() {
	_, err := s.db.Query(`INSERT INTO days(Date, ElevenAM, Noon, OnePM, TwoPM, ThreePM, FourPM, FivePM, SixPM, SevenPM, EightPM, NinePM, TenPM, ElevenPM, Total) VALUES('1990-01-02',10,10,10,10,10,10,10,10,10,10,10,10,10,10)`)
	if err != nil {
		s.T().Fatal(err)
	}

	days, err := s.store.GetDays()
	if err != nil {
		s.T().Fatal(err)
	}
	nDays := len(days)
	if nDays != 1 {
		s.T().Errorf("Inncorect count, wanted 1, got %d", nDays)
	}

	expectedDay := Day{"1990-01-02T00:00:00Z",10,10,10,10,10,10,10,10,10,10,10,10,10,10}
	if *days[0] != expectedDay {
		s.T().Errorf("Incorrect details, expected %v, got %v", expectedDay, *days[0])
	}
}



func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Get(mockServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}
}
func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Post(mockServer.URL+"/hello", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}
func TestRouter(t *testing.T) {
	r :=newRouter()

	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL +"/hello")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := "Hello World!"
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}
func TestHandler(t *testing.T){
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(handler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `Hello World!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",actual, expected)
	}
}
