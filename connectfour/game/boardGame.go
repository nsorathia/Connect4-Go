package game

import (
	"games/connectfour/config"
	"games/connectfour/board"
	"games/connectfour/dataDevice"
	"games/connectfour/enums"
	"games/connectfour/player"
	"strconv"
	"strings"
)

func init() {
	NewGame = NewBoardGame
}

//BoardGame represents a game with two human players
type BoardGame struct {
	device  dataDevice.DataDevice
	board   board.Board
	players []player.Player
}

//NewBoardGame sets the player's names and returns a Two Player Game
func NewBoardGame() Game {

	device := dataDevice.NewDataDevice()
	board := board.NewBoard()
	players := setUpPlayers(device)

	tpg := BoardGame{
		device:  device,
		board:   board,
		players: players,
	}

	return &tpg
}

func setUpPlayers(device dataDevice.DataDevice) []player.Player {
	var player2 player.Player
	var player1 player.Player

	device.Write("Enter Player1's name: ")
	player1Name := device.Read()
	humanPlayer1 := player.NewHumanPlayer(player1Name, enums.Red, device)
	player1 = &humanPlayer1

	device.Write("Would you like to player against the computer? Y/N ")

	if isOnePlayerGame := strings.ToUpper(device.Read()); isOnePlayerGame == "Y" {
		device.Write("What level of difficulty would you like the computer to play? 1...5   1:Easy ... 5:Master")
		difficultyLevel, err := strconv.Atoi(device.Read())
		if err != nil {
			difficultyLevel = config.GetInt("difficulty-level")
		}

		//TODO  validate user input
		aiPlayer := player.NewAIPlayer(difficultyLevel)
		player2 = &aiPlayer

	} else {
		device.Write("Enter Player2's name: ")
		player2Name := device.Read()
		humanPlayer2 := player.NewHumanPlayer(player2Name, enums.Yellow, device)
		player2 = &humanPlayer2
	}

	return []player.Player{player1, player2}
}

//Players returns the game's players
func (g *BoardGame) Players() []player.Player {
	return g.players
}

//Device returns the dataDevice
func (g *BoardGame) Device() dataDevice.DataDevice {
	return g.device
}

//Board returns the game Board
func (g *BoardGame) Board() board.Board {
	return g.board
}
