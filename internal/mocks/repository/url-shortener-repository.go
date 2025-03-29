package mocks

import "github.com/stretchr/testify/mock"

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) StoreShortURL(shortURL, longURL string) error {
	args := m.Called(shortURL, longURL)
	return args.Error(0)
}

func (m *MockRepo) GetLongURL(shortURL string) (string, error) {
	args := m.Called(shortURL)
	return args.String(0), args.Error(1)
}

func (m *MockRepo) GetShortURL(longURL string) string {
	args := m.Called(longURL)
	return args.String(0)
}
