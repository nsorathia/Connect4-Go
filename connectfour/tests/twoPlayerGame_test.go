package tests

import (
	"games/connectfour/game"
	"github.com/stretchr/testify/assert"
	"testing"
	//"games/connectfour/player"
)

func TestNewTwoPlayerGameReturnsAGame(t *testing.T) {

	newGame := game.NewConnectFourGame()
	assert.NotNil(t, newGame)
	assert.NotNil(t, newGame.Board())
	assert.NotNil(t, newGame.Device())
	assert.NotNil(t, newGame.Players)

	//Two players
	assert.NotNil(t, len(newGame.Players()) == 2)
	assert.NotNil(t, newGame.Players()[0].Token() != newGame.Players()[1].Token())
}
