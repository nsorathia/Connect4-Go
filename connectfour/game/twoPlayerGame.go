package game

import (
	"games/connectfour/enums"
	"games/connectfour/board"
	"games/connectfour/player"
	"games/connectfour/dataDevice"
)


func init() {
	NewGame = NewTwoPlayerGame
}

//TwoPlayerGame represents a game with two human players 
type TwoPlayerGame struct {
	device dataDevice.DataDevice
	board board.Board
	players []player.Player
}

//NewTwoPlayerGame sets the player's names and returns a Two Player Game 
func  NewTwoPlayerGame() Game {

	device := dataDevice.NewDataDevice()
	board := board.NewBoard()
	players := setUpTwoPlayers(device)
	
	tpg := TwoPlayerGame{
		device: device,
		board: board,
		players: players,
	}

	return &tpg
}

func setUpTwoPlayers(device dataDevice.DataDevice) []player.Player {
	
	device.Write("Enter Player1's name: ")
	player1Name := device.Read()
	player1 := player.NewHumanPlayer(player1Name, enums.Red, device)

	device.Write("Enter Player2's name: ") 
	player2Name := device.Read()
	player2 := player.NewHumanPlayer(player2Name, enums.Yellow, device)	

	players:= []player.Player{ &player1, &player2 }
	return players
}


//Players returns the game's players
func (g *TwoPlayerGame) Players() []player.Player {
	return g.players
}

//Device returns the dataDevice 
func (g *TwoPlayerGame) Device() dataDevice.DataDevice {
	return g.device
}

//Board returns the game Board
func (g *TwoPlayerGame) Board() board.Board {
	return g.board
}