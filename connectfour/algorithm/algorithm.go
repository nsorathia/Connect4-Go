package algorithm

import (
	"games/connectfour/board"
	"games/connectfour/enums"
)

type Algorithm interface {
	CalculateBestMove(board.Board, enums.Token) int
}

type factory func() Algorithm

var NewAlgorithm factory
