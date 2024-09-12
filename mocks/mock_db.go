package mocks

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Select(dest interface{}, query string, args ...interface{}) error {
	argsMock := m.Called(dest, query, args)
	return argsMock.Error(0)
}

// Implementação mock para o método Get
func (m *MockDB) Get(dest interface{}, query string, args ...interface{}) error {
    argsMock := m.Called(dest, query, args)
    return argsMock.Error(0)
}

// Implementação mock para o método Exec (usado para Update e Delete)
func (m *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
    argsMock := m.Called(query, args)
    return nil, argsMock.Error(1)
}

// Implementação mock para o método QueryRow (usado para Insert)
func (m *MockDB) QueryRow(query string, args ...interface{}) *sql.Row {
    argsMock := m.Called(query, args)
    return argsMock.Get(0).(*sql.Row)
}