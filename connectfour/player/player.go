package player

import (
	"games/connectfour/board"
	"games/connectfour/enums"
)

//Player is an abstract representation on a ConnectFour Player
type Player interface {
	Move(board.Board) (int, error)
	Name() string
	Token() enums.Token
}

