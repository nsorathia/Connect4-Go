package game

import (
	"fmt"
	"games/connectfour/board"
	"games/connectfour/data"
	"games/connectfour/dataDevice"
	"games/connectfour/player"
)

//Game represents a game between players
type Game interface {
	GameID() int
	Players() []player.Player
	Device() dataDevice.DataDevice
	Board() board.Board
}

type factory func(data.Repository, dataDevice.DataDevice) Game

//NewGame is a factory method which returns a game type
var NewGame factory

//Play facilitates the game
func Play() {

	repo := data.NewRepository()
	device := dataDevice.NewDataDevice()

	game := NewGame(repo, device)
	board := game.Board()
	players := game.Players()
	displayBoard(board, device)

	gameHasWinner := false
	totalMoves := board.Columns() * board.Rows()
	for i := 0; i < totalMoves; i++ {

		player := getPlayer(i, players)
		move, _ := player.Move(board)

		repo.SaveMove(game.GameID(), player.Id(), move)

		board.SetPlayerMove(move, player.Token())
		displayBoard(board, device)

		if gameHasWinner = board.IsWin(); gameHasWinner {
			device.Write(fmt.Sprintf("%v - YOU WON!!!", player.Name()))
			break
		}
	}

	if !gameHasWinner {
		device.Write("IT's a TIE!!")
	}
}

func getPlayer(move int, players []player.Player) player.Player {
	if move%2 == 0 {
		return players[0]
	}
	return players[1]
}

func displayBoard(board board.Board, device dataDevice.DataDevice) {
	device.Write(board.ToString())
	device.Write("\n")
}
