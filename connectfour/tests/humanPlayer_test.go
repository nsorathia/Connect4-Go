package tests

import (
	"games/connectfour/enums"
	"games/connectfour/player"
	"games/connectfour/tests/mocks"
	"games/connectfour/utilities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveReturnsErrorIfBoardIsNil(t *testing.T) {

	mockDevice := mocks.MockDataDevice{}
	human := player.NewHumanPlayer(0, -1, "Frodo", enums.Red, &mockDevice)

	_, err := human.Move(nil)
	assert.Error(t, err)
}

func TestMoveCallsGetAvailableMoves(t *testing.T) {

	testUserChoice := "4"
	availableMoves := []int{0, 3, 4, 5}

	mockDevice := mocks.MockDataDevice{}
	mockDevice.Mock.On("Read").Return(testUserChoice)
	mockDevice.Mock.On("Write").Return()

	mockBoard := mocks.MockBoard{}
	mockBoard.Mock.On("GetAvailableMoves").Return(availableMoves)

	human := player.NewHumanPlayer(0, -1, "Frodo", enums.Red, &mockDevice)
	_, _ = human.Move(&mockBoard)

	mockBoard.Mock.AssertCalled(t, "GetAvailableMoves")
}

func TestMoveEnforcesValidBoardChoice(t *testing.T) {

	//Move enforces the user's choice is contained in the board's available moves.
	testUserChoice := "4"
	availableMoves := []int{0, 3, 4, 5}

	mockDevice := mocks.MockDataDevice{}
	mockDevice.Mock.On("Read").Return(testUserChoice)
	mockDevice.Mock.On("Write").Return()

	mockBoard := mocks.MockBoard{}
	mockBoard.Mock.On("GetAvailableMoves").Return(availableMoves)

	human := player.NewHumanPlayer(0, -1, "Frodo", enums.Red, &mockDevice)
	choice, err := human.Move(&mockBoard)

	assert.Nil(t, err)
	assert.True(t, utility.Contains(availableMoves, choice))
}
