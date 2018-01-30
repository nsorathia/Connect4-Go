package main

import (
	_ "games/connectfour/board"
	_ "games/connectfour/dataDevice"
	"games/connectfour/enums"
	"games/connectfour/player"
	"games/abstract"
)

func main() {
	board := abstract.NewBoard()	
	device := abstract.NewDataDevice()

	player1 := player.NewHumanPlayer("Frodo", enums.Red, device)
	player1.Move(board)

}
