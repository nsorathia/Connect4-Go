package board

import (
	"math"
)

//BoardVariant represents a particular version of the board when an token is dropped
type BoardVersion struct {
	Board    Board
	Versions []BoardVersion
	Score    int
}

//NewBoardVariant returns a new BoardVariant with default values
func NewBoardVersion(board Board, versions []BoardVersion) BoardVersion {
	return BoardVersion{
		Board:    board,
		Versions: versions,
		Score:    math.MinInt64,
	}
}
