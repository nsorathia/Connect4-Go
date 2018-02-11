package board

import (
	"bytes"
	"errors"
	"fmt"
	"games/connectfour/enums"
	"games/connectfour/utilities"
)

func init() {
	if GameType == "tictactoe" {
		NewBoard = NewTicTacToeBoard
	}
}

const ConsecutiveTokensForTicTacToeWin = 3

type TicTacToeBoard struct {
	TTTGrid                [][]enums.Token
	TTTLastMoveColumnIndex int //Column Index of Last Player's Move
}

func NewTicTacToeBoard() Board {

	newBoard, err := CreateNewTicTacToeBoard(3, 3)
	if err != nil {
		return nil
	}

	return &newBoard
}

func CreateNewTicTacToeBoard(rows, columns int) (TicTacToeBoard, error) {
	if rows < 3 {
		return TicTacToeBoard{}, errors.New("C4Board needs at least 4 rows")
	}

	if columns < 3 {
		return TicTacToeBoard{}, errors.New("C4Board needs at least 4 columns")
	}

	grid := make([][]enums.Token, rows)
	for i := range grid {
		grid[i] = make([]enums.Token, columns)
	}

	return TicTacToeBoard{
		TTTGrid:                grid,
		TTTLastMoveColumnIndex: -1,
	}, nil
}

func (b *TicTacToeBoard) TicTacRows() int {
	return len(b.TTTGrid)
}

func (b *TicTacToeBoard) TicTacColumns() int {
	return len(b.TTTGrid[0])
}

func (b *TicTacToeBoard) Clone() Board {

	duplicate := make([][]enums.Token, len(b.TTTGrid))
	for i := range b.TTTGrid {
		duplicate[i] = make([]enums.Token, len(b.TTTGrid[i]))
		copy(duplicate[i], b.TTTGrid[i])
	}

	clone := TicTacToeBoard{
		TTTGrid:                duplicate,
		TTTLastMoveColumnIndex: b.TTTLastMoveColumnIndex,
	}

	return &clone
}

func (b *TicTacToeBoard) SetPlayerMove(columnNumber int, token enums.Token) error {

	if columnNumber < 1 || columnNumber > b.TicTacColumns() {
		return errors.New("Column move must be within bounds of the board")
	}

	if token == enums.Empty {
		return errors.New("Empty token can not be used to set player's move")
	}

	moveIndex := columnNumber - 1

	if b.isColumnFull(moveIndex) {
		return fmt.Errorf("Column: %v is full", columnNumber)
	}

	for i := b.TicTacRows() - 1; i >= 0; i-- {
		if b.TTTGrid[i][moveIndex] == enums.Empty {
			b.TTTGrid[i][moveIndex] = token
			b.TTTLastMoveColumnIndex = moveIndex
			break
		}
	}

	return nil
}

func (b *TicTacToeBoard) GetAvailableMoves() []int {

	availableColumns := []int{}
	columnLength := b.TicTacColumns()

	for i := 0; i < columnLength; i++ {
		if !b.isColumnFull(i) {
			availableColumns = append(availableColumns, i)
		}
	}

	return availableColumns
}

func (b *TicTacToeBoard) IsWin() bool {

	lastMoveRowIndex := b.lastMoveRowIndex()
	lastMoveColumnIndex := b.TTTLastMoveColumnIndex

	if b.isHorizontalWin(lastMoveRowIndex, lastMoveColumnIndex) ||
		b.isVerticalWin(lastMoveRowIndex, lastMoveColumnIndex) ||
		b.isDownDiagonalWin(lastMoveRowIndex, lastMoveColumnIndex) ||
		b.isUpDiagonalWin(lastMoveRowIndex, lastMoveColumnIndex) {
		return true
	}

	return false
}

//ToString is a method with returns a string representation of the C4 board
func (b *TicTacToeBoard) ToString() string {

	buffer := bytes.Buffer{}

	for rows := range b.TTTGrid {
		for columns := range b.TTTGrid[rows] {
			var marker string

			switch token := b.TTTGrid[rows][columns]; token {
			case enums.Red:
				marker = " X "
			case enums.Yellow:
				marker = " O "
			default:
				marker = " # "
			}

			buffer.WriteString(marker)
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func (b *TicTacToeBoard) isColumnFull(column int) bool {
	isFull := false
	if b.TTTGrid[0][column] != enums.Empty {
		isFull = true
	}

	return isFull
}

func (b *TicTacToeBoard) lastMoveRowIndex() int {

	numOfRows := b.TicTacRows()

	i := 0
	for ; i < numOfRows; i++ {
		if b.TTTGrid[i][b.TTTLastMoveColumnIndex] != enums.Empty {
			break
		}
	}

	return i
}

func (b *TicTacToeBoard) isHorizontalWin(rowIndex, columnIndex int) bool {
	leftBound := utility.Max(0, columnIndex-3)
	rightBound := utility.Min(b.TicTacColumns()-1, columnIndex+3)
	token := b.TTTGrid[rowIndex][columnIndex]
	count := 0

	for i := leftBound; i <= rightBound; i++ {
		if b.TTTGrid[rowIndex][i] == token {
			count++
			if count == ConsecutiveTokensForTicTacToeWin {
				return true
			}
		} else {
			count = 0
		}
	}

	return false
}

func (b *TicTacToeBoard) isVerticalWin(rowIndex, columnIndex int) bool {
	lowBound := rowIndex
	highBound := utility.Min(rowIndex+3, b.TicTacRows()-1)
	token := b.TTTGrid[rowIndex][columnIndex]
	count := 0

	//no need to go any further if the
	if highBound-lowBound < ConsecutiveTokensForTicTacToeWin-1 {
		return false
	}

	for i := lowBound; i <= highBound; i++ {
		if b.TTTGrid[i][columnIndex] == token {
			count++
			if count == ConsecutiveTokensForTicTacToeWin {
				return true
			}
		} else {
			count = 0
		}
	}

	return false
}

func (b *TicTacToeBoard) isDownDiagonalWin(rowIndex, columnIndex int) bool {
	// X
	//   X
	//	   X
	//       X
	//set the left-upper bound and right-lower bound
	leftUpperRowIndex, leftUpperColumnIndex := b.leftUpperCoordinate(rowIndex, columnIndex)
	rightLowerRowIndex, rightLowerColumnIndex := b.rightLowerCoordinate(rowIndex, columnIndex)

	//no need to seach further if there are less than 4 tokens to check
	if rightLowerRowIndex-leftUpperRowIndex < ConsecutiveTokensForTicTacToeWin-1 {
		return false
	}

	count := 0
	token := b.TTTGrid[rowIndex][columnIndex]
	i, j := leftUpperRowIndex, leftUpperColumnIndex

	for i <= rightLowerRowIndex && j <= rightLowerColumnIndex {
		if b.TTTGrid[i][j] == token {
			count++
			if count == ConsecutiveTokensForTicTacToeWin {
				return true
			}
		} else {
			count = 0
		}
		i++
		j++
	}

	return false
}

func (b *TicTacToeBoard) isUpDiagonalWin(rowIndex, columnIndex int) bool {
	//       X
	//	   X
	//   X
	// X
	//set the left-lower and right-upper bounds
	leftLowerRowIndex, leftLowerColumnIndex := b.leftLowerCoordinate(rowIndex, columnIndex)
	rightUpperRowIndex, rightUpperColumnIndex := b.rightUpperCoordinate(rowIndex, columnIndex)

	//test if that there are less than 4 tokens otherwise no need to go further
	if rightUpperColumnIndex-leftLowerColumnIndex < ConsecutiveTokensForTicTacToeWin-1 {
		return false
	}

	count := 0
	token := b.TTTGrid[rowIndex][columnIndex]
	i, j := leftLowerRowIndex, leftLowerColumnIndex

	for i >= rightUpperRowIndex && j <= rightUpperColumnIndex {
		if b.TTTGrid[i][j] == token {
			count++
			if count == ConsecutiveTokensForTicTacToeWin {
				return true
			}
		} else {
			count = 0
		}

		i--
		j++
	}

	return false
}

func (b *TicTacToeBoard) leftLowerCoordinate(rowIndex, columnIndex int) (int, int) {

	rowUpperBound := b.TicTacRows() - 1

	switch {
	case rowIndex == rowUpperBound || columnIndex == 0:
		return rowIndex, columnIndex

	case rowIndex == rowUpperBound-1 || columnIndex == 1:
		return rowIndex + 1, columnIndex - 1

	case rowIndex == rowUpperBound-2 || columnIndex == 2:
		return rowIndex + 2, columnIndex - 2

	default: //case rowIndex <= rowUpperBound - 3 || columnIndex >= 3:
		return rowIndex + 3, columnIndex - 3
	}
}

func (b *TicTacToeBoard) rightUpperCoordinate(rowIndex, columnIndex int) (int, int) {

	columnUpperBound := b.TicTacColumns() - 1

	switch {
	case rowIndex == 0 || columnIndex == columnUpperBound:
		return rowIndex, columnIndex

	case rowIndex == 1 || columnIndex == columnUpperBound-1:
		return rowIndex - 1, columnIndex + 1

	case rowIndex == 2 || columnIndex == columnUpperBound-2:
		return rowIndex - 2, columnIndex + 2

	default: //case rowIndex >= 3 || columnIndex <= columnUpperBound - 3:
		return rowIndex - 3, columnIndex + 3
	}
}

//leftUpperCoordinate returns the coordinate of the left-lower cell to evaluate for a down diagnal win
func (b *TicTacToeBoard) leftUpperCoordinate(rowIndex, columnIndex int) (int, int) {
	switch {
	case rowIndex == 0 || columnIndex == 0:
		return rowIndex, columnIndex

	case rowIndex == 1 || columnIndex == 1:
		return rowIndex - 1, columnIndex - 1

	case rowIndex == 2 || columnIndex == 2:
		return rowIndex - 2, columnIndex - 2

	default: //case rowIndex >= 3 || columnIndex >= 3:
		return rowIndex - 3, columnIndex - 3
	}
}

func (b *TicTacToeBoard) rightLowerCoordinate(rowIndex, columnIndex int) (int, int) {

	rowUpperBound := b.TicTacRows() - 1
	columnUpperBound := b.TicTacColumns() - 1

	switch {
	case rowIndex == rowUpperBound || columnIndex == columnUpperBound:
		return rowIndex, columnIndex

	case rowIndex == rowUpperBound-1 || columnIndex == columnUpperBound-1:
		return rowIndex + 1, columnIndex + 1

	case rowIndex == rowUpperBound-2 || columnIndex == columnUpperBound-2:
		return rowIndex + 2, columnIndex + 2

	default: //case rowIndex <= rowUpperBound - 3 || columnIndex <= columnUpperBound - 3:
		return rowIndex + 3, columnIndex + 3
	}
}

func (b *TicTacToeBoard) Grid() [][]enums.Token {
	return b.TTTGrid
}

func (b *TicTacToeBoard) LastMove() int {
	return b.TTTLastMoveColumnIndex
}

func (b *TicTacToeBoard) Rows() int {
	return len(b.TTTGrid)
}

func (b *TicTacToeBoard) Columns() int {
	return len(b.TTTGrid[0])
}
