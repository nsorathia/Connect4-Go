package data

import (
	"games/connectfour/enums"
)

//Repository is abstraction to a underlying data repository like a database or file store
type Repository interface {
	SaveGame(gameType string) (int, error)
	SavePlayer(gameID int, name string, token enums.Token) (int, error)
	SaveMove(gameID, playerID, move int)
}

type factory func() Repository

//NewRepository is a factory that will return a Implemented Repository
var NewRepository factory


