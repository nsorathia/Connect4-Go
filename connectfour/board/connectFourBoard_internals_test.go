package board

import (
	"fmt"
	"games/connectfour/enums"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestIsColumnFullReturnsTrueIfColumnIsFull(t *testing.T) {
	c4 := createBoard(6, 7)

	//populate the first column
	for i := range c4.Grid {
		c4.Grid[i][0] = enums.Red
	}

	//test first column
	assert.True(t, c4.isColumnFull(0))

	//test second column
	assert.False(t, c4.isColumnFull(1))
}

func TestLastRowReturnsRowIndexOfLastMove(t *testing.T) {
	c4 := createBoard(6, 1)
	randomLastRowIndex := random(0, 5)

	//populate the first with random number of Tokens
	for i := 5; i >= randomLastRowIndex; i-- {
		c4.Grid[i][0] = enums.Red
	}
	fmt.Println(c4.ToString())

	c4.LastMoveColumnIndex = 1
	assert.EqualValues(t, randomLastRowIndex, c4.lastMoveRowIndex())
}

func TestIsHorizontalWinReturnsTrueIfLastMoveProducesHorizontalWin(t *testing.T) {

	var c4 ConnectFourBoard
	var lb, rb int

	// X X X X # # #
	//create and populate board
	c4 = createBoard(1, 7)
	lb, rb = 0, 3
	for i := lb; i <= rb; i++ {
		c4.Grid[0][i] = enums.Red
	}
	//test each token if last dropped
	for i := lb; i <= rb; i++ {
		assert.True(t, c4.isHorizontalWin(0, i))
	}

	// # # # X X X X
	c4 = createBoard(1, 7)
	lb, rb = 3, 6
	for i := lb; i <= rb; i++ {
		c4.Grid[0][i] = enums.Red
	}
	//test each token if last dropped
	for i := lb; i <= rb; i++ {
		assert.True(t, c4.isHorizontalWin(0, i))
	}

	// # # X X X X #
	c4 = createBoard(1, 7)
	lb, rb = 2, 5
	for i := lb; i <= rb; i++ {
		c4.Grid[0][i] = enums.Red
	}
	//test each token if last dropped
	for i := lb; i <= rb; i++ {
		assert.True(t, c4.isHorizontalWin(0, i))
	}

	// X # X X X # #
	c4 = createBoard(1, 7)
	lb, rb = 2, 4
	for i := lb; i <= rb; i++ {
		c4.Grid[0][i] = enums.Red
	}
	c4.Grid[0][0] = enums.Red
	//test each token if last dropped
	for i := lb; i <= rb; i++ {
		assert.False(t, c4.isHorizontalWin(0, i))
	}
}

func TestIsVerticalWinReturnsTrueIfLastMoveProducesVerticalWin(t *testing.T) {

	var c4 ConnectFourBoard
	var lb, hb int

	// X
	// X
	// X
	// X
	// O
	// O
	c4 = createBoard(6, 1)
	lb, hb = 0, 3
	for i := lb; i <= hb; i++ {
		c4.Grid[i][0] = enums.Red
	}
	//test last dropped token
	assert.True(t, c4.isVerticalWin(lb, 0))

	// #
	// #
	// X
	// X
	// X
	// X
	c4 = createBoard(6, 1)
	lb, hb = 2, 5
	for i := lb; i <= hb; i++ {
		c4.Grid[i][0] = enums.Red
	}
	//test last dropped token
	assert.True(t, c4.isVerticalWin(lb, 0))

	// #
	// X
	// X
	// X
	// X
	// O
	c4 = createBoard(6, 1)
	lb, hb = 1, 4
	for i := lb; i <= hb; i++ {
		c4.Grid[i][0] = enums.Red
	}
	//test last dropped token
	assert.True(t, c4.isVerticalWin(lb, 0))

	// #
	// X
	// X
	// O
	// X
	// O
	c4 = createBoard(6, 1)
	c4.Grid[1][0] = enums.Red
	c4.Grid[2][0] = enums.Red
	c4.Grid[3][0] = enums.Yellow
	c4.Grid[4][0] = enums.Red
	c4.Grid[5][0] = enums.Yellow

	//test last dropped token
	assert.False(t, c4.isVerticalWin(1, 0))

}

func TestIsDownDiagonalWinReturnsTrueIfLastMoveProducesADownDiagonalWin(t *testing.T) {
	var c4 ConnectFourBoard

	// # # # # X # #
	// # # # # # X #
	// # # # # # # X
	// # # # # # # #
	// # # # # # # #
	// # # # # # # #
	c4 = createBoard(6,7)
	c4.Grid[0][4] = enums.Red
	c4.Grid[1][5] = enums.Red
	c4.Grid[2][6] = enums.Red

	fmt.Println(c4.ToString())

	assert.False(t, c4.isDownDiagonalWin(0, 4))
	assert.False(t, c4.isDownDiagonalWin(1, 5))
	assert.False(t, c4.isDownDiagonalWin(2, 6))
	

	// # # # # # # #
	// # # X # # # #
	// # # # X # # #
	// # # # # X # #
	// # # # # # X #
	// # # # # # # #
	c4 = createBoard(6,7)
	c4.Grid[1][2] = enums.Red
	c4.Grid[2][3] = enums.Red
	c4.Grid[3][4] = enums.Red
	c4.Grid[4][5] = enums.Red

	fmt.Println(c4.ToString())

	assert.True(t, c4.isDownDiagonalWin(1, 2))
	assert.True(t, c4.isDownDiagonalWin(3, 4))
	assert.True(t, c4.isDownDiagonalWin(4, 5))
	assert.True(t, c4.isDownDiagonalWin(2, 3))

	// # # # # # # #
	// # # # # # # #
	// # # # # # # #
	// # # # # # # #
	// X # # # # # #
	// # X # # # # #
	c4 = createBoard(6,7)
	c4.Grid[4][0] = enums.Red
	c4.Grid[5][1] = enums.Red

	fmt.Println(c4.ToString())

	assert.False(t, c4.isDownDiagonalWin(5, 1))
	assert.False(t, c4.isDownDiagonalWin(4, 0))	
}

func TestIsUpDiagonalWinReturnsTrueIfLastMoveProducesAUpDiagonalWin(t *testing.T) {
	var c4 ConnectFourBoard

	// # # X # # # #
	// # X # # # # #
	// X # # # # # #
	// # # # # # # #
	// # # # # # # #
	// # # # # # # #
	c4 = createBoard(6,7)
	c4.Grid[0][2] = enums.Red
	c4.Grid[1][1] = enums.Red
	c4.Grid[2][0] = enums.Red

	fmt.Println(c4.ToString())

	assert.False(t, c4.isUpDiagonalWin(0, 2))
	assert.False(t, c4.isUpDiagonalWin(1, 1))
	assert.False(t, c4.isUpDiagonalWin(2, 0))
	

	// # # # # # # #
	// # # # X # # #
	// # # X # # # #
	// # X # # # # #
	// X # # # # # #
	// # # # # # # #
	c4 = createBoard(6,7)
	c4.Grid[4][0] = enums.Red
	c4.Grid[3][1] = enums.Red
	c4.Grid[2][2] = enums.Red
	c4.Grid[1][3] = enums.Red

	fmt.Println(c4.ToString())

	assert.True(t, c4.isUpDiagonalWin(4, 0))
	assert.True(t, c4.isUpDiagonalWin(3, 1))
	assert.True(t, c4.isUpDiagonalWin(2, 2))
	assert.True(t, c4.isUpDiagonalWin(1, 3))

	// # # # # # # #
	// # # # # # # #
	// # # # # # # #
	// # # # # # # #
	// # # # # # # X
	// # # # # # X #
	c4 = createBoard(6,7)
	c4.Grid[4][6] = enums.Red
	c4.Grid[5][5] = enums.Red

	fmt.Println(c4.ToString())

	assert.False(t, c4.isUpDiagonalWin(5, 5))
	assert.False(t, c4.isUpDiagonalWin(4, 6))	
}



//helper test functions
func createBoard(rows, columns int) ConnectFourBoard {
	grid := make([][]enums.Token, rows)
	for i := range grid {
		grid[i] = make([]enums.Token, columns)
	}
	return ConnectFourBoard{grid, -1}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
