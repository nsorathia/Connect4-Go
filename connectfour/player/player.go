package player

import (
	"games/connectfour/board"
	"games/connectfour/enums"
)


//Player is an abstract representation on a CopnnectFour Player
type Player interface {
	Move(board.Board) int
	Name() string
	Token() enums.Token
}
