package mocks

import (
	"games/connectfour/enums"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) SaveGame(gameType string) (int, error) {
	args := m.Called(gameType)
	return args.Int(0), args.Error(1)
}

func (m *MockRepository) SavePlayer(gameID int, name string, token enums.Token) (int, error) {
	args := m.Called(gameID, name, token)
	return args.Int(0), args.Error(1)
}

func (m *MockRepository) SaveMove(gameID, playerID, move int) {
	m.Called(gameID, playerID, move)
}