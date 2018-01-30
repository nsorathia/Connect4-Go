package tests

import (
	"fmt"
	"games/connectfour/board"
	"games/connectfour/enums"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestNewBoardToReturnInitalizedBoardWith6x7dimensions(t *testing.T) {

	actual := board.NewConnectFourBoard()

	assert.Equal(t, 6, len(actual.Grid()))
	assert.Equal(t, 7, len(actual.Grid()[0]))
}

func TestCreateNewBoardToReturnInitailizedBoard(t *testing.T) {

	var table = []struct {
		row int // row
		col int // column
	}{
		{3, 4},
		{4, 4},
		{6, 7},
		{4, 2},
		{11, 15},
	}

	for _, data := range table {

		board, err := board.CreateNewBoard(data.row, data.col)
		if err != nil {
			emsg := err.Error()
			if emsg != "C4Board needs at least 4 rows" && emsg != "C4Board needs at least 4 columns" {

				assert.Equal(t, data.row, len(board.C4Grid))    //row
				assert.Equal(t, data.col, len(board.C4Grid[0])) //column
				assert.Equal(t, -1, board.C4LastMoveColumnIndex)
			}
		}
	}
}

func TestRowsToReturnNumberOfGridRows(t *testing.T) {
	var table = []struct {
		c4   board.ConnectFourBoard
		rows int
	}{
		{createBoard(3, 5), 3},
		{createBoard(6, 3), 6},
		{createBoard(4, 5), 4},
		{createBoard(7, 9), 7},
		{createBoard(8, 5), 8},
	}

	for _, td := range table {
		assert.Equal(t, td.rows, td.c4.C4Rows())
	}
}

func TestColumnsToReturnNumberOfGridColumns(t *testing.T) {
	var table = []struct {
		c4      board.ConnectFourBoard
		columns int
	}{
		{createBoard(3, 5), 5},
		{createBoard(6, 3), 3},
		{createBoard(4, 5), 5},
		{createBoard(7, 9), 9},
		{createBoard(8, 5), 5},
	}

	for _, td := range table {
		assert.Equal(t, td.columns, td.c4.C4Columns())
	}
}

func TestToStringReturnsATypeOfString(t *testing.T) {

	c4 := createBoard(6, 7)
	assert.Equal(t, "string", reflect.TypeOf(c4.ToString()).String())
}

// //REDO:XXXXXXXX
func TestCloneReturnsACopyOfC4Board(t *testing.T) {
	c1 := createBoard(6, 7)
	c1.C4Grid[3][4] = enums.Red

	clone := c1.Clone()

	//test length and width of clone
	assert.True(t, len(clone.Grid()) == len(c1.C4Grid))
	assert.True(t, len(clone.Grid()[0]) == len(c1.C4Grid[0]))

	//Ceck each cell
	cloneGrid := clone.Grid()
	for i := 0; i < clone.Rows(); i++ {
		for j := 0; j < clone.Columns(); j++ {
			assert.True(t, cloneGrid[i][j] == c1.C4Grid[i][j])
		}
	}

	//modify c1
	c1.C4LastMoveColumnIndex = 4
	c1.C4Grid[2][2] = enums.Yellow

	// //check that objects are not equal ater c1 is modified
	assert.False(t, cloneGrid[2][2] == c1.C4Grid[2][2])
	assert.False(t, clone.LastMove() == c1.C4LastMoveColumnIndex)
}

func TestSetPlayerMoveReturnsErrorIfMoveIsOutOfRange(t *testing.T) {

	row := random(4, 8)
	column := random(4, 8)

	c4 := createBoard(row, column)

	err := c4.SetPlayerMove(-1, enums.Red)
	assert.NotNil(t, err)

	err = c4.SetPlayerMove(column+1, enums.Red)
	assert.NotNil(t, err)

	err = c4.SetPlayerMove(2, enums.Red)
	assert.Nil(t, err)

}

func TestSetPlayerMoveReturnsErrorIfTokenIsEmpty(t *testing.T) {
	row := random(4, 8)
	column := random(4, 8)

	c4 := createBoard(row, column)

	err := c4.SetPlayerMove(2, enums.Empty)
	assert.NotNil(t, err)
}

func TestSetPlayerMoveReturnsErrorIfcolumnIsFull(t *testing.T) {
	row := random(4, 8)
	column := random(4, 8)

	//populate the first column of the board
	c4 := createBoard(row, column)
	for i := range c4.C4Grid {
		c4.C4Grid[i][0] = enums.Red
	}

	//try to add another token
	err := c4.SetPlayerMove(0, enums.Yellow)
	assert.NotNil(t, err)
}

func TestSetPlayerMoveReturnsNilIfColumnIsNotFull(t *testing.T) {
	row := random(4, 8)
	column := random(4, 8)

	//populate the first 3 rows of column 1
	c4 := createBoard(row, column)
	bottomRow := len(c4.C4Grid)
	c4.C4Grid[bottomRow-1][0] = enums.Red
	c4.C4Grid[bottomRow-2][0] = enums.Red
	c4.C4Grid[bottomRow-3][0] = enums.Red

	//try to add another token to column 0
	fmt.Println(c4.ToString())

	err := c4.SetPlayerMove(1, enums.Yellow)
	assert.Nil(t, err)
}

func TestSetPlayerMoveSetsTokenInNextAvailableRow(t *testing.T) {
	row := random(4, 8)
	column := random(4, 8)

	//populate the first 3 rows of column 0
	c4 := createBoard(row, column)
	bottomRow := len(c4.C4Grid)
	c4.C4Grid[bottomRow-1][2] = enums.Red
	c4.C4Grid[bottomRow-2][2] = enums.Red
	c4.C4Grid[bottomRow-3][2] = enums.Red

	//test 4th row before SetMove
	testBefore := (c4.C4Grid[bottomRow-4][2] == enums.Empty)
	assert.True(t, testBefore)

	//Set token in 4th row
	fmt.Println(c4.ToString())
	c4.SetPlayerMove(3, enums.Yellow)
	fmt.Println(c4.ToString())

	//test 4th row after SetMove
	testAfter := (c4.C4Grid[bottomRow-4][2] == enums.Yellow)
	assert.True(t, testAfter)
}

func TestSetPlayerMoveUpdatesLastMoveField(t *testing.T) {
	row := random(4, 8)
	column := random(4, 8)

	//populate the first 3 rows of column 0
	c4 := createBoard(row, column)
	assert.True(t, c4.C4LastMoveColumnIndex == -1)

	move := c4.C4Columns() - 2

	c4.SetPlayerMove(move, enums.Red)
	assert.True(t, c4.C4LastMoveColumnIndex == move-1)

}

func TestGetAvailableMovesReturnsSliceOfAvailableColumnIndexes(t *testing.T) {
	row := 6
	column := 7

	//populate columns 2 and 5
	c4 := createBoard(row, column)
	c4.C4Grid[0][1] = enums.Red
	c4.C4Grid[0][4] = enums.Yellow

	expected := []int{0, 2, 3, 5, 6}
	actual := c4.GetAvailableMoves()

	assert.EqualValues(t, expected, actual)
}

func TestIsLastMoveWinIfBoardHasAWin(t *testing.T) {
	c4 := createBoard(6, 7)

	// # # # # # # #
	// # # # # O # #
	// # # # O # # #
	// # # O # # # #
	// # O # # # # #
	// # O # # # # #
	c4.C4Grid[1][4] = enums.Yellow
	c4.C4Grid[2][3] = enums.Yellow
	c4.C4Grid[3][2] = enums.Yellow
	c4.C4Grid[4][1] = enums.Yellow
	c4.C4Grid[5][1] = enums.Yellow
	c4.C4LastMoveColumnIndex = 1
	fmt.Println(c4.ToString())

	assert.True(t, c4.IsWin())

	// # # # O # # #
	// # # # # O # #
	// # # # # # X #
	// # # # # # # O
	// # # # # # # X
	// # # # # # # X
	c4.C4Grid[0][3] = enums.Yellow
	c4.C4Grid[1][4] = enums.Yellow
	c4.C4Grid[2][5] = enums.Red
	c4.C4Grid[3][6] = enums.Yellow
	c4.C4Grid[4][6] = enums.Red
	c4.C4Grid[5][6] = enums.Red
	c4.C4LastMoveColumnIndex = 6
	fmt.Println(c4.ToString())

	assert.False(t, c4.IsWin())

	// # # # # # # #
	// # # # # # # #
	// # # # # # # #
	// # # # # # # #
	// # # # # # # #
	// # # O O O O #
	c4.C4Grid[5][2] = enums.Yellow
	c4.C4Grid[5][3] = enums.Yellow
	c4.C4Grid[5][4] = enums.Yellow
	c4.C4Grid[5][5] = enums.Yellow
	c4.C4LastMoveColumnIndex = 4
	fmt.Println(c4.ToString())

	assert.True(t, c4.IsWin())

	// # # # # # # #
	// # # # # # # #
	// # # X # # # #
	// # # X # # # #
	// # # X # # # #
	// # # X # # # #
	c4.C4Grid[2][2] = enums.Red
	c4.C4Grid[3][2] = enums.Red
	c4.C4Grid[4][2] = enums.Red
	c4.C4Grid[5][2] = enums.Red
	c4.C4LastMoveColumnIndex = 2
	fmt.Println(c4.ToString())

	assert.True(t, c4.IsWin())

}

func TestConnectFourBoardImplementsBoard(t *testing.T) {
	c4 := createBoard(6, 7)

	c4.Clone()

}

func createBoard(rows, columns int) board.ConnectFourBoard {
	grid := make([][]enums.Token, rows)
	for i := range grid {
		grid[i] = make([]enums.Token, columns)
	}

	return board.ConnectFourBoard{
		C4Grid:                grid,
		C4LastMoveColumnIndex: -1,
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
