package algorithm

import (
	"fmt"
	"games/connectfour/board"
	"games/connectfour/enums"
	"games/connectfour/utilities"
	"math"
	"os"
	"strconv"
	"games/connectfour/config"
)

func init() {
	temp := config.GetString("algorithm")
	if temp == "minmax" {
		NewAlgorithm = NewMinMaxAlgorithm
	}
}

//MinMaxAlgorithm is a recursive algorithm for choosing a players next best move:  https://en.wikipedia.org/wiki/Minimax
type MinMaxAlgorithm struct {
	difficultyLevel int
}

//NewMinMaxAlgorithm returns an instance of MinMaxAlgorithm.
func NewMinMaxAlgorithm() Algorithm {

	level, err := strconv.Atoi(os.Getenv("DIFFICULTY_LEVEL"))
	if err != nil {
		level = 5
	}

	return &MinMaxAlgorithm{difficultyLevel: level}
}

//CalculateBestMove utilizes the minMaxAlgorithm to determine the best move.
func (m *MinMaxAlgorithm) CalculateBestMove(gameBoard board.Board, token enums.Token) int {

	versionGraph := new(board.BoardVersion)

	versionGraph.Board = gameBoard

	graphVariants(versionGraph, token, m.difficultyLevel)

	scores := make([]int, 0)
	for _, variant := range versionGraph.Versions {
		scores = append(scores, calculateScore(&variant, m.difficultyLevel))
	}

	bestMove := chooseBestMove(scores)

	return bestMove + 1 //return columnNumber not columnIndex
}

func graphVariants(variant *board.BoardVersion, token enums.Token, level int) {

	if level == 0 {
		return
	}

	moveVersions := variant.Board.GetAvailableMoves()

	for i := 0; i < len(moveVersions); i++ {

		clone := variant.Board.Clone()

		clone.SetPlayerMove(moveVersions[i]+1, token)

		newVariant := board.NewBoardVersion(clone, nil)

		graphVariants(&newVariant, getOpposingToken(token), level-1)

		variant.Versions = append(variant.Versions, newVariant)
	}
}

func calculateScore(boardVersion *board.BoardVersion, level int) int {

	score := 0

	points := int(math.Pow10(level))

	if boardVersion.Board.IsWin() {
		score = points

	} else if opponentNextMoveCanWin(boardVersion) {
		score = -points

	} else {

		//check opposite
		for _, opponentsVariant := range boardVersion.Versions {
			score -= calculateScore(&opponentsVariant, level-1)
		}
	}

	//debug(boardVersion, score, level)

	return score
}

func chooseBestMove(scores []int) int {

	numberOfMoves := len(scores)
	bestScore := math.MinInt64
	bestMoves := make([]int, 0)

	for i := 0; i < numberOfMoves; i++ {
		score := scores[i]
		if score > bestScore {
			bestScore = score
			bestMoves = make([]int, 0)
			bestMoves = append(bestMoves, i)

		} else if score == bestScore {
			bestMoves = append(bestMoves, i)
		}
	}

	//take random move from slice
	randomIndex := 0
	if len(bestMoves) > 1 {
		randomIndex = utility.Random(0, len(bestMoves)-1)
	}
	return bestMoves[randomIndex]
}

func opponentNextMoveCanWin(variant *board.BoardVersion) bool {
	for _, oppMove := range variant.Versions {
		if oppMove.Board.IsWin() {
			return true
		}
	}
	return false
}

func getOpposingToken(token enums.Token) enums.Token {
	if token == enums.Red {
		return enums.Yellow
	}
	return enums.Red
}

func debug(boardVersion *board.BoardVersion, score, level int) {

	if level > 5 {
		fmt.Printf("Level: %v   Score: %v \n", level, score)
		fmt.Println(boardVersion.Board.ToString())
	}
}
