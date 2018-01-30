package main

import (
	"games/connectfour/board"
	"games/connectfour/dataDevice"
	"games/connectfour/enums"
	"games/connectfour/player"
)

func main() {
	board := board.NewBoard()	
	device := dataDevice.NewDataDevice()

	player1 := player.NewHumanPlayer("Frodo", enums.Red, device)
	player1.Move(board)
}
