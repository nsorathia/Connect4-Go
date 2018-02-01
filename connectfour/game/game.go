package game

import (
	"fmt"
	"games/connectfour/board"
	"games/connectfour/dataDevice"
	"games/connectfour/player"
)

type Game interface {
	Players() []player.Player
	Device() dataDevice.DataDevice
	Board() board.Board
}

type factory func() Game
var NewGame factory


func Play() {

	game := NewGame()
	board := game.Board()
	players := game.Players()
	device := game.Device()

	displayBoard(board, device)

	gameHasWinner := false
	totalMoves := board.Columns() * board.Rows()
	for i := 0; i < totalMoves; i++ {
		
		player := getPlayer(i, players)
		move, _ := player.Move(board)
		board.SetPlayerMove(move, player.Token())
		displayBoard(board, device)

		if gameHasWinner = board.IsWin(); gameHasWinner {
			device.Write(fmt.Sprintf("%v - YOU WON!!!", player.Name()))
			break;
		} 
	}

	if !gameHasWinner {
		device.Write("IT's a TIE!!")
	}
}

func getPlayer(move int, players []player.Player) player.Player {
	if (move % 2 == 0) {
		return players[0]
	} 
	return players[1]
}

func displayBoard(board board.Board, device dataDevice.DataDevice) {
	device.Write(board.ToString())
	device.Write("\n");
}
