package mocks

import (
	"games/connectfour/board"
	"games/connectfour/enums"
	"github.com/stretchr/testify/mock"
)

type MockPlayer struct {
	mock.Mock
}

func (m *MockPlayer) Name() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockPlayer) Token() enums.Token {
	args := m.Called()
	return args.Get(0).(enums.Token)
}

func (m *MockPlayer) Move(board.Board) (int, error) {
	args := m.Called()
	return args.Int(0), args.Get(0).(error)
}