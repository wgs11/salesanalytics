package main

import "github.com/stretchr/testify/mock"

type MockStore struct {
	mock.Mock
}


func (m *MockStore) GetEmployee(id int)([]*Employee, error) {
	rets := m.Called()
	return rets.Get(0).([]*Employee), rets.Error(1)
}
func (m *MockStore) CreateDay(day *Day) error {
	rets := m.Called(day)
	return rets.Error(0)
}

func (m *MockStore) GetDay(estimate float32)([]*Day, error) {
	rets := m.Called()
	return rets.Get(0).([]*Day), rets.Error(1)
}

func (m *MockStore) GetDays()([]*Day, error) {
	rets := m.Called()
	return rets.Get(0).([]*Day), rets.Error(1)
}

func (m *MockStore) GetPlace(location string)(*Place, error) {
	rets := m.Called()
	return rets.Get(0).(*Place), rets.Error(1)
}

func (m *MockStore) CreateUser(creds *Credentials) error {
	rets := m.Called(creds)
	return rets.Error(0)
}
func (m *MockStore) CheckUser(creds *Credentials) error {
	rets := m.Called(creds)
	return rets.Error(0)
}

func (m *MockStore) GetReviews(location string) ([]*Review, error) {
	rets := m.Called(location)
	return rets.Get(0).([]*Review), rets.Error(1)
}

func (m *MockStore) GetReview(location string, date string) (*Review, error) {
	rets := m.Called(location, date)
	return rets.Get(0).(*Review), rets.Error(1)
}

func InitMockStore() *MockStore {
	s := new(MockStore)
	store = s
	return s
}
