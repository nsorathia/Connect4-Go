package abstract

import (
	"games/connectfour/enums"
)

//Board is an abstract representation of the ConnectFourBoard
type Board interface {
	Grid() [][]enums.Token
	Rows() int
	Columns() int
	LastMove() int

	Clone() Board
	IsWin() bool
	SetPlayerMove(number int, token enums.Token) error
	GetAvailableMoves() []int
	ToString() string
}

type factory func() Board

var NewBoard factory
