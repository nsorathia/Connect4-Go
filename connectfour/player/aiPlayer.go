package player

import (
	"strconv"
	"os"
	"games/connectfour/algorithm"
	"games/connectfour/board"
	"games/connectfour/enums"
)

//AIPlayer is an Artificial Intelligent player that utilizes to MinMax algorithm to calculate its best move
type AIPlayer struct {
	name string
	token enums.Token
	algo algorithm.Algorithm
}

func NewAIPlayer(difficultyLevel int) AIPlayer {

	os.Setenv("DIFFICULTY_LEVEL", strconv.Itoa(difficultyLevel))

	aiPlayer := AIPlayer{
		name: "R2D2",
		token: enums.Yellow,
		algo: algorithm.NewAlgorithm(),
	}

	return aiPlayer
}

func (a *AIPlayer) Name() string {
	return a.name
}

func (a *AIPlayer) Token() enums.Token {
	return a.token
}

func (a *AIPlayer) Move(board board.Board) (int, error) {
	return a.algo.CalculateBestMove(board, a.token), nil
} 

