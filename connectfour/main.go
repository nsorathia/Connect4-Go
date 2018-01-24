package main

import (
	"fmt"
	"games/connectfour/board"
)

func main() {

	c4, err := board.CreateNewBoard(6, 7)
	if err != nil {
		return
	}
	
	fmt.Print(c4.ToString())
}
