package board

import (
	"games/connectfour/enums"
	"games/connectfour/config"
)

//Board is an abstract representation of the ConnectFourBoard
type Board interface {
	Grid() [][]enums.Token
	Rows() int
	Columns() int
	LastMove() int

	Clone() Board
	IsWin() bool
	SetPlayerMove(columnNumber int, token enums.Token) error
	GetAvailableMoves() []int
	ToString() string
}


var GameType = config.GetString("game")

var NewBoard factory

type factory func() Board
