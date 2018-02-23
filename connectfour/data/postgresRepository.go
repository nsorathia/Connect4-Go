package data

import (
	"database/sql"
	"games/connectfour/config"
	"games/connectfour/enums"
	"strconv"

	_ "github.com/lib/pq" //the postgres driver utilized by database package above
)

const InsertGame = "INSERT INTO games(type) VALUES($1) RETURNING id"
const InsertPlayer = "INSERT INTO players(name, game_id, token) VALUES($1, $2, $3) RETURNING id"
const InsertMove = "INSERT INTO game_play(game_id, player_id, move) VALUES($1, $2, $3)"

var connectionString = config.GetString("dbconn")

func NewPostgresRepository() Repository {
	return &PostgresRepository{}
}

type PostgresRepository struct {
}

func (pr *PostgresRepository) SaveGame(gameType string) (int, error) {

	if gameType == "connectfour" || gameType == "tictactoe" {

		db := openDB()
		defer close(db)
		return saveGame(db, gameType)
	}

	panic("gameType must be set to either connectfour or tictactoe")
}

func (pr *PostgresRepository) SavePlayer(gameID int, name string, token enums.Token) (int, error) {
	db := openDB()
	defer close(db)
	return savePlayer(db, gameID, name, token), nil
}

func (pr *PostgresRepository) SaveMove(gameID, playerID, move int) {

	db := openDB()
	defer close(db)
	saveMove(db, gameID, playerID, move)
}

func saveMove(db *sql.DB, gameID, playerID, move int) {
	_, err := db.Exec(InsertMove, strconv.Itoa(gameID), strconv.Itoa(playerID), strconv.Itoa(move))
	if err != nil {
		panic(err)
	}
}

func saveGame(db *sql.DB, gameName string) (int, error) {

	lastInsertID := 0
	err := db.QueryRow(InsertGame, gameName).Scan(&lastInsertID)
	return lastInsertID, err
}

func savePlayer(db *sql.DB, gameID int, name string, token enums.Token) int {
	tokenStr := "1"
	if token == enums.Yellow {
		tokenStr = "2"
	}

	lastInsertID := 0
	err := db.QueryRow(InsertPlayer, name, strconv.Itoa(gameID), tokenStr).Scan(&lastInsertID)

	if err != nil {
		panic(err)
	}

	return lastInsertID
}

func openDB() *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	return db
}

func close(db *sql.DB) {
	dbErr := db.Close()
	if dbErr != nil {
		panic(dbErr)
	}
}
