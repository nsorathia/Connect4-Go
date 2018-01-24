package player

import (
	"games/connectfour/enums"
)

//HumanPlayer represents a Human ConnectFourPlayer
type HumanPlayer struct {
	Name string
	Token enums.Token
}