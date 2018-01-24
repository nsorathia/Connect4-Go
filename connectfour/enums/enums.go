package enums

//token represents the Player's game piece for a ConnectFour game
type Token int

//oken 
const (
	Empty Token = iota //0   Represents the Tokens on a ConnectFour Board
	Red                 //1
	Yellow              //2
)