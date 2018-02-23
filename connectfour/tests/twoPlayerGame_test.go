package tests

import (
	"games/connectfour/tests/mocks"
	"games/connectfour/game"
	"github.com/stretchr/testify/assert"
	"testing"
	"games/connectfour/enums"
	
)

func TestNewTwoPlayerGameReturnsAGame(t *testing.T) {

	mockDevice := mocks.MockDataDevice{}
	mockDevice.Mock.On("Write")
	mockDevice.Mock.On("Read").Return("Frodo")

	mockRepo := mocks.MockRepository{}
	mockRepo.Mock.On("SaveGame", "connectfour").Return(0, nil)
	mockRepo.Mock.On("SavePlayer", 0, "Frodo", enums.Red).Return(-1, nil)
	mockRepo.Mock.On("SavePlayer", 0, "Frodo", enums.Yellow).Return(-2, nil)
	mockRepo.Mock.On("SaveMove", 0, -1, 4, )

	newGame := game.NewBoardGame(&mockRepo, &mockDevice)
	assert.NotNil(t, newGame)
	assert.NotNil(t, newGame.Board())
	assert.NotNil(t, newGame.Device())
	assert.NotNil(t, newGame.Players)

	//Two players
	assert.NotNil(t, len(newGame.Players()) == 2)
	assert.NotNil(t, newGame.Players()[0].Token() != newGame.Players()[1].Token())
}
