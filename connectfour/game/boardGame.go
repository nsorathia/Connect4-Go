package game

import (
	"games/connectfour/board"
	"games/connectfour/config"
	"games/connectfour/data"
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
	gameId  int
	device  dataDevice.DataDevice
	board   board.Board
	players []player.Player
}

//NewBoardGame sets the player's names and returns a Two Player Game
func NewBoardGame(repo data.Repository, device dataDevice.DataDevice) Game {

	gameName := config.GetString("game")
	gameID, _ := repo.SaveGame(gameName)

	board := board.NewBoard()
	players := setUpPlayers(gameID, device, repo)

	return &BoardGame{
		gameId:  gameID,
		device:  device,
		board:   board,
		players: players,
	}

}

//START HERE pass game ID to players and save to db
func setUpPlayers(gameId int, device dataDevice.DataDevice, repo data.Repository) []player.Player {
	var player2 player.Player
	var player1 player.Player

	device.Write("Enter Player1's name: ")
	player1Name := device.Read()
	player1ID, _ := repo.SavePlayer(gameId, player1Name, enums.Red)

	humanPlayer1 := player.NewHumanPlayer(gameId, player1ID, player1Name, enums.Red, device)
	player1 = &humanPlayer1

	device.Write("Would you like to player against the computer? Y/N ")

	if isOnePlayerGame := strings.ToUpper(device.Read()); isOnePlayerGame == "Y" {
		device.Write("What level of difficulty would you like the computer to play? 1...5   1:Easy ... 5:Master")
		difficultyLevel, err := strconv.Atoi(device.Read())
		if err != nil {
			difficultyLevel = config.GetInt("difficulty-level")
		}

		player2Id, _ := repo.SavePlayer(gameId, "R2D2", enums.Yellow)
		aiPlayer := player.NewAIPlayer(gameId, player2Id, difficultyLevel)
		player2 = &aiPlayer

	} else {
		device.Write("Enter Player2's name: ")
		player2Name := device.Read()
		player2Id, _ := repo.SavePlayer(gameId, player1Name, enums.Yellow)

		humanPlayer2 := player.NewHumanPlayer(gameId, player2Id, player2Name, enums.Yellow, device)
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

//GameId represent the unique recorded id of the game
func (g *BoardGame) GameID() int {
	return g.gameId
}
