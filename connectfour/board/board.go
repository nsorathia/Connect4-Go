package board

import (
	"games/connectfour/config"
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
	SetPlayerMove(columnNumber int, token enums.Token) error
	GetAvailableMoves() []int
	ToString() string
}

type factory func() Board

//NewBoard is a a factory that abstracts the implementation of a game board.  
var NewBoard factory

func init() {
	gameType := config.GetString("game")

	switch gameType {

	case "tictactoe":
		NewBoard = NewTicTacToeBoard

	default:
		NewBoard = NewConnectFourBoard
	}
}
