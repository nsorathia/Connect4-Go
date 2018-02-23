package data

import (
	"games/connectfour/enums"
	"games/connectfour/utilities"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"strconv"
	"testing"
)

func TestSaveGamePanicsIfGameTypeIsNotSet(t *testing.T) {
	repo := PostgresRepository{}
	assert.Panics(t, func() { repo.SaveGame("") })
	assert.Panics(t, func() { repo.SaveGame("monopoly") })
}

func TestSaveGameReturnsGameID(t *testing.T) {
	db, sqlMock, _ := sqlmock.New()
	defer db.Close()

	columns := []string{"id"}

	gameName := "connectfour"

	sqlMock.ExpectQuery("INSERT INTO games").
		WithArgs(gameName).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("1"))

	saveGame(db, gameName)
	assert.Nil(t, sqlMock.ExpectationsWereMet())
}

func TestSavePlayerReturnsPlayerId(t *testing.T) {
	db, sqlMock, _ := sqlmock.New()
	defer db.Close()

	columns := []string{"id"}

	name := "Frodo"
	gameID := utility.Random(1, 100)
	token := enums.Yellow

	sqlMock.ExpectQuery("INSERT INTO players").
		WithArgs(name, strconv.Itoa(gameID), "2").
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("1"))

	savePlayer(db, gameID, name, token)
	assert.Nil(t, sqlMock.ExpectationsWereMet())
}

func TestSaveMoveInsertsRecordInGamePlay(t *testing.T) {
	db, sqlMock, _ := sqlmock.New()
	defer db.Close()

	gameID := utility.Random(1, 100)
	playerID := utility.Random(1, 100)
	move := utility.Random(1, 7)

	sqlMock.ExpectExec("INSERT INTO game_play").
		WithArgs(strconv.Itoa(gameID), strconv.Itoa(playerID), strconv.Itoa(move)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	saveMove(db, gameID, playerID, move)
	assert.Nil(t, sqlMock.ExpectationsWereMet())
}
