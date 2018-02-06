package algorithm

import (
	"fmt"
	"games/connectfour/board"
	"games/connectfour/enums"
	"games/connectfour/utilities"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestGetOpposingTokenReturnsOpposingColorToken(t *testing.T) {

	token := getOpposingToken(enums.Yellow)
	assert.True(t, token == enums.Red)

	token = getOpposingToken(enums.Red)
	assert.True(t, token == enums.Yellow)
}

func TestGraphVariantUpdatesVariant_N_LevelsDeep(t *testing.T) {

	token := enums.Red
	c4 := createBoard(6, 7)
	c4.SetPlayerMove(4, token)
	variant := new(board.BoardVersion)
	variant.Board = &c4

	randomLevel := random(3, 4)

	//test
	graphVariants(variant, token, randomLevel)

	testVariant := variant.Versions[0]

	for i := 0; i < randomLevel-1; i++ {
		testVariant = testVariant.Versions[0]
		assert.NotNil(t, testVariant.Board)
	}
}

func TestCalculateScoreReturnsScoreForEachMove(t *testing.T) {

	level := 6
	game, token := createRandomMovePlayBoard()

	fmt.Println(game.ToString())

	variant := new(board.BoardVersion)
	variant.Board = game

	//im lazy
	graphVariants(variant, token, level)

	for _, moveVariant := range variant.Versions {
		moveVariant.Score = calculateScore(&moveVariant, level)
		assert.True(t, moveVariant.Score != math.MinInt64)
	}
}

func TestCalculateBestMove(t *testing.T) {

	gameboard, token := createRandomMovePlayBoard()

	fmt.Println(gameboard.ToString())

	algo := NewMinMaxAlgorithm()

	bestMove := algo.CalculateBestMove(gameboard, getOpposingToken(token))

	assert.True(t, utility.Contains([]int{0, 1, 2, 3, 4, 5}, bestMove))
}

func createBoard(rows, columns int) board.ConnectFourBoard {
	grid := make([][]enums.Token, rows)
	for i := range grid {
		grid[i] = make([]enums.Token, columns)
	}

	return board.ConnectFourBoard{
		C4Grid:                grid,
		C4LastMoveColumnIndex: -1,
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func createRandomMovePlayBoard() (board.Board, enums.Token) {
	game := board.NewBoard()
	randomMoveCount := 6 //each player plays three so we dont have a winner
	token := enums.Empty
	for i := 0; i < randomMoveCount; i++ {

		if token == enums.Red {
			token = enums.Yellow
		} else {
			token = enums.Red
		}

		randomMove := random(1, 7)
		game.SetPlayerMove(randomMove, token)
	}

	return game, token
}
