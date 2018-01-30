package abstract

import (
	"games/connectfour/enums"
)

//Player is an abstract representation on a ConnectFour Player
type Player interface {
	Move(Board) (int, error)
	Name() string
	Token() enums.Token
}

