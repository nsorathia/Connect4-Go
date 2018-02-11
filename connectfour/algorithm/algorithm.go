package algorithm

import (
	"games/connectfour/config"
	"games/connectfour/board"
	"games/connectfour/enums"
)

type Algorithm interface {
	CalculateBestMove(board.Board, enums.Token) int
}

type factory func() Algorithm

var NewAlgorithm factory

var AlgorithmType = config.GetString("algorithm")
