package tests

import (
	"time"
	"fmt"
	"games/connectfour/board"
	"games/connectfour/enums"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"testing"
)

func TestNewBoardToReturnInitalizedBoardWith6x7dimensions(t *testing.T) {
	actual, err := board.NewBoard()
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, 6, len(actual.Grid))
	assert.Equal(t, 7, len(actual.Grid[0]))
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

				assert.Equal(t, data.row, len(board.Grid))    //row
				assert.Equal(t, data.col, len(board.Grid[0])) //column
				assert.Equal(t, -1, board.LastMoveColumnIndex)
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
		assert.Equal(t, td.rows, td.c4.Rows())
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
		assert.Equal(t, td.columns, td.c4.Columns())
	}
}

func TestToStringReturnsATypeOfString(t *testing.T) {

	c4 := createBoard(6, 7)
	assert.Equal(t, "string", reflect.TypeOf(c4.ToString()).String())
}

func TestCloneReturnsACopyOfC4Board(t *testing.T) {
	rows := random(4, 9)
	columns := random(4, 9)

	c1 := createBoard(rows, columns)
	c1.Grid[3][4] = enums.Red

	clone := c1.Clone()
	//check structs are equal after clone
	assert.EqualValues(t, c1, clone)

	c1.LastMoveColumnIndex = 4
	c1.Grid[2][2] = enums.Yellow

	//check that objects are not equal ater c1 is modified
	assert.False(t, assert.ObjectsAreEqual(c1, clone))
}

func TestSetMoveReturnsErrorIfMoveIsOutOfRange(t *testing.T) {

	row := random(4, 8)
	column := random(4, 8)

	c4 := createBoard(row, column)

	err := c4.DropPlayerToken(-1, enums.Red)
	assert.NotNil(t, err)

	err = c4.DropPlayerToken(column+1, enums.Red)
	assert.NotNil(t, err)

	err = c4.DropPlayerToken(2, enums.Red)
	assert.Nil(t, err)

}

func TestSetMoveReturnsErrorIfTokenIsEmpty(t *testing.T) {
	row := random(4, 8)
	column := random(4, 8)

	c4 := createBoard(row, column)

	err := c4.DropPlayerToken(2, enums.Empty)
	assert.NotNil(t, err)
}

func TestSetMoveReturnsErrorIfcolumnIsFull(t *testing.T) {
	row := random(4, 8)
	column := random(4, 8)

	//populate the first column of the board
	c4 := createBoard(row, column)
	for i := range c4.Grid {
		c4.Grid[i][0] = enums.Red
	}

	//try to add another token
	err := c4.DropPlayerToken(0, enums.Yellow)
	assert.NotNil(t, err)
}

func TestSetMoveReturnsNilIfColumnIsNotFull(t *testing.T) {
	row := random(4, 8)
	column := random(4, 8)

	//populate the first 3 rows of column 1
	c4 := createBoard(row, column)
	bottomRow := len(c4.Grid)
	c4.Grid[bottomRow-1][0] = enums.Red
	c4.Grid[bottomRow-2][0] = enums.Red
	c4.Grid[bottomRow-3][0] = enums.Red

	//try to add another token to column 0
	fmt.Println(c4.ToString())

	err := c4.DropPlayerToken(1, enums.Yellow)
	assert.Nil(t, err)
}

func TestSetMoveSetsTokenInNextAvailableRow(t *testing.T) {
	row := random(4, 8)
	column := random(4, 8)

	//populate the first 3 rows of column 0
	c4 := createBoard(row, column)
	bottomRow := len(c4.Grid)
	c4.Grid[bottomRow-1][2] = enums.Red
	c4.Grid[bottomRow-2][2] = enums.Red
	c4.Grid[bottomRow-3][2] = enums.Red

	//test 4th row before SetMove
	testBefore := (c4.Grid[bottomRow-4][2] == enums.Empty)
	assert.True(t, testBefore)

	//Set token in 4th row
	fmt.Println(c4.ToString())
	c4.DropPlayerToken(3, enums.Yellow)
	fmt.Println(c4.ToString())

	//test 4th row after SetMove
	testAfter := (c4.Grid[bottomRow-4][2] == enums.Yellow)
	assert.True(t, testAfter)
}

func TestSetMoveUpdatesLastMoveField(t *testing.T) {
	row := random(4, 8)
	column := random(4, 8)

	//populate the first 3 rows of column 0
	c4 := createBoard(row, column)
	assert.True(t, c4.LastMoveColumnIndex == -1)

	move := c4.Columns() - 2

	c4.DropPlayerToken(move, enums.Red)
	assert.True(t, c4.LastMoveColumnIndex == move-1)

}

func TestGetAvailableMovesReturnsSliceOfAvailableColumnIndexes(t *testing.T) {
	row := 6
	column := 7

	//populate columns 2 and 5
	c4 := createBoard(row, column)
	c4.Grid[0][1] = enums.Red
	c4.Grid[0][4] = enums.Yellow

	expected := []int{0, 2, 3, 5, 6}
	actual := c4.GetAvailableMoves()

	assert.EqualValues(t, expected, actual)
}

func TestIsLastMoveWinIfBoardHasAWin(t *testing.T) {
	c4 := createBoard(6,7)

	// # # # # # # #
	// # # # # O # #
	// # # # O # # #
	// # # O # # # #
	// # O # # # # #
	// # O # # # # #
	c4.Grid[1][4] = enums.Yellow
	c4.Grid[2][3] = enums.Yellow
	c4.Grid[3][2] = enums.Yellow
	c4.Grid[4][1] = enums.Yellow
	c4.Grid[5][1] = enums.Yellow
	c4.LastMoveColumnIndex = 1
	fmt.Println(c4.ToString())

	assert.True(t, c4.IsLastMoveWin())

	// # # # O # # #
	// # # # # O # #
	// # # # # # X #
	// # # # # # # O
	// # # # # # # X
	// # # # # # # X
	c4.Grid[0][3] = enums.Yellow
	c4.Grid[1][4] = enums.Yellow
	c4.Grid[2][5] = enums.Red
	c4.Grid[3][6] = enums.Yellow
	c4.Grid[4][6] = enums.Red
	c4.Grid[5][6] = enums.Red
	c4.LastMoveColumnIndex = 6
	fmt.Println(c4.ToString())
	
	assert.False(t, c4.IsLastMoveWin())

	// # # # # # # #
	// # # # # # # #
	// # # # # # # #
	// # # # # # # #
	// # # # # # # #
	// # # O O O O #
	c4.Grid[5][2] = enums.Yellow
	c4.Grid[5][3] = enums.Yellow
	c4.Grid[5][4] = enums.Yellow
	c4.Grid[5][5] = enums.Yellow
	c4.LastMoveColumnIndex = 4
	fmt.Println(c4.ToString())
	
	assert.True(t, c4.IsLastMoveWin())

	// # # # # # # #
	// # # # # # # #
	// # # X # # # #
	// # # X # # # #
	// # # X # # # #
	// # # X # # # #
	c4.Grid[2][2] = enums.Red
	c4.Grid[3][2] = enums.Red
	c4.Grid[4][2] = enums.Red
	c4.Grid[5][2] = enums.Red
	c4.LastMoveColumnIndex = 2
	fmt.Println(c4.ToString())
	
	assert.True(t, c4.IsLastMoveWin())
	
}

func createBoard(rows, columns int) board.ConnectFourBoard {
	grid := make([][]enums.Token, rows)
	for i := range grid {
		grid[i] = make([]enums.Token, columns)
	}
	return board.ConnectFourBoard{grid, -1}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
