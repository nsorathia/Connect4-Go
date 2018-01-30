package mocks

import (
	"games/abstract"
	"games/connectfour/enums"
	"github.com/stretchr/testify/mock"
)

type MockBoard struct {
	mock.Mock
}

func NewBoard() abstract.Board {
	return &MockBoard{}
}

func (b *MockBoard) Clone() abstract.Board {
	args := b.Called()
	return  args.Get(0).(*MockBoard)
}

func (b *MockBoard) SetPlayerMove(columnNumber int, token enums.Token) error {
	args := b.Called(columnNumber, token)
	return  args.Error(0)
}

func (b *MockBoard) GetAvailableMoves() []int {
	args := b.Called()
	return args.Get(0).([]int)
}

func (b *MockBoard) IsWin() bool {
	return b.Called().Bool(0)
}

func (b *MockBoard) ToString() string {
	args := b.Called()
	return args.String(0)
}

func (b *MockBoard) Grid() [][]enums.Token {
	args := b.Called()
	return args.Get(0).([][]enums.Token)
}

func (b *MockBoard) LastMove() int {
	args := b.Called()
	return args.Int(0)
}

func (b *MockBoard) Rows() int {
	args := b.Called()
	return args.Int(0)
}

func (b *MockBoard) Columns() int {
	args := b.Called()
	return args.Int(0)
}
