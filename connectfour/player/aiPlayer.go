package player

import (
	"games/connectfour/algorithm"
	"games/connectfour/board"
	"games/connectfour/enums"
	"os"
	"strconv"
)

const AIPlayerName = "R2D2"

//AIPlayer is an Artificial Intelligent player that utilizes to MinMax algorithm to calculate its best move
type AIPlayer struct {
	id    int
	name  string
	token enums.Token
	algo  algorithm.Algorithm
}

func NewAIPlayer(gameId, id, difficultyLevel int) AIPlayer {

	os.Setenv("DIFFICULTY_LEVEL", strconv.Itoa(difficultyLevel))
	
	aiPlayer := AIPlayer{
		id:    id,
		name:  AIPlayerName,
		token: enums.Yellow,
		algo:  algorithm.NewAlgorithm(),
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

func (a *AIPlayer) Id() int {
	return a.id
}
