package board

import (
	"games/connectfour/enums"
)

//Board is an abstract representation of the ConnectFourBoard
type Board interface {
	IsLastMoveWin() bool
	Clone() Board
	DropPlayerToken(number int, token enums.Token)
	GetAvailableMoves() []int
	ToString()
}
