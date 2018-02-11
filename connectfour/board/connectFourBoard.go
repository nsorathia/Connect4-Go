package board

import (
	"bytes"
	"errors"
	"fmt"
	"games/connectfour/enums"
	"games/connectfour/utilities"
	"games/connectfour/config"
)

func init() {
	gametype := config.GetString("game")
	if gametype == "connectfour" {
		NewBoard = NewConnectFourBoard
	}
}

//ConsecutiveTokensForWin is the number of adjacent tokens fom a single player to claim a win
const ConsecutiveTokensForWin = 4

//ConnectFourBoard represent the game board
type ConnectFourBoard struct {
	C4Grid                [][]enums.Token
	C4LastMoveColumnIndex int //Column Index of Last Player's Move
}

//NewConnectFourBoard creates a ConnectFourBoard with default dimensions:  6 rows and 7 columns
func NewConnectFourBoard() Board {

	newBoard, err := CreateNewBoard(6, 7)
	if err != nil {
		return nil
	}

	return &newBoard
}

//CreateNewBoard creates a ConnectFourBoard given the dimensions
func CreateNewBoard(rows, columns int) (ConnectFourBoard, error) {
	if rows < 4 {
		return ConnectFourBoard{}, errors.New("C4Board needs at least 4 rows")
	}

	if columns < 4 {
		return ConnectFourBoard{}, errors.New("C4Board needs at least 4 columns")
	}

	grid := make([][]enums.Token, rows)
	for i := range grid {
		grid[i] = make([]enums.Token, columns)
	}

	return ConnectFourBoard{
		C4Grid:                grid,
		C4LastMoveColumnIndex: -1,
	}, nil
}

//C4Rows is a method that returns the number of rows in the ConnectFourBoard
func (b *ConnectFourBoard) C4Rows() int {
	return len(b.C4Grid)
}

//C4Columns is a method that returns the number of C4Columns for the board.
func (b *ConnectFourBoard) C4Columns() int {
	return len(b.C4Grid[0])
}

//Clone is a method whihc returns a deep copy of the C4 board
func (b *ConnectFourBoard) Clone() Board {

	duplicate := make([][]enums.Token, len(b.C4Grid))
	for i := range b.C4Grid {
		duplicate[i] = make([]enums.Token, len(b.C4Grid[i]))
		copy(duplicate[i], b.C4Grid[i])
	}

	clone := ConnectFourBoard{
		C4Grid:                duplicate,
		C4LastMoveColumnIndex: b.C4LastMoveColumnIndex,
	}

	return &clone
}

//SetPlayerMove is a method that updates a C4 Board with a player's token given teh Column Number
func (b *ConnectFourBoard) SetPlayerMove(columnNumber int, token enums.Token) error {

	if columnNumber < 1 || columnNumber > b.C4Columns() {
		return errors.New("Column move must be within bounds of the board")
	}

	if token == enums.Empty {
		return errors.New("Empty token can not be used to set player's move")
	}

	moveIndex := columnNumber - 1

	if b.isColumnFull(moveIndex) {
		return fmt.Errorf("Column: %v is full", columnNumber)
	}

	for i := b.C4Rows() - 1; i >= 0; i-- {
		if b.C4Grid[i][moveIndex] == enums.Empty {
			b.C4Grid[i][moveIndex] = token
			b.C4LastMoveColumnIndex = moveIndex
			break
		}
	}

	return nil
}

//GetAvailableMoves is a method that returns a slice contianing the index of non-full columns
func (b *ConnectFourBoard) GetAvailableMoves() []int {

	availableColumns := []int{}
	columnLength := b.C4Columns()

	for i := 0; i < columnLength; i++ {
		if !b.isColumnFull(i) {
			availableColumns = append(availableColumns, i)
		}
	}

	return availableColumns
}

//IsWin is a method which determines if the last players Move was a win
func (b *ConnectFourBoard) IsWin() bool {

	lastMoveRowIndex := b.lastMoveRowIndex()
	lastMoveColumnIndex := b.C4LastMoveColumnIndex

	if b.isHorizontalWin(lastMoveRowIndex, lastMoveColumnIndex) ||
		b.isVerticalWin(lastMoveRowIndex, lastMoveColumnIndex) ||
		b.isDownDiagonalWin(lastMoveRowIndex, lastMoveColumnIndex) ||
		b.isUpDiagonalWin(lastMoveRowIndex, lastMoveColumnIndex) {
		return true
	}

	return false
}

//ToString is a method with returns a string representation of the C4 board
func (b *ConnectFourBoard) ToString() string {

	buffer := bytes.Buffer{}

	for rows := range b.C4Grid {
		for columns := range b.C4Grid[rows] {
			var marker string

			switch token := b.C4Grid[rows][columns]; token {
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

//isColumnFull is a method which determines if a column is full
func (b *ConnectFourBoard) isColumnFull(column int) bool {
	isFull := false
	if b.C4Grid[0][column] != enums.Empty {
		isFull = true
	}

	return isFull
}

//lastMoveRowIndex is a method which returns the rowIndex of the last move
func (b *ConnectFourBoard) lastMoveRowIndex() int {

	numOfRows := b.C4Rows()

	i := 0
	for ; i < numOfRows; i++ {
		if b.C4Grid[i][b.C4LastMoveColumnIndex] != enums.Empty {
			break
		}
	}

	return i
}

//isHorizontalWin is a method with determines if the last move was part of a horizontal Win
func (b *ConnectFourBoard) isHorizontalWin(rowIndex, columnIndex int) bool {
	leftBound := utility.Max(0, columnIndex-3)
	rightBound := utility.Min(b.C4Columns()-1, columnIndex+3)
	token := b.C4Grid[rowIndex][columnIndex]
	count := 0

	for i := leftBound; i <= rightBound; i++ {
		if b.C4Grid[rowIndex][i] == token {
			count++
			if count == ConsecutiveTokensForWin {
				return true
			}
		} else {
			count = 0
		}
	}

	return false
}

//isVerticalWin is a method which determines if the last move was part of a vertical win
func (b *ConnectFourBoard) isVerticalWin(rowIndex, columnIndex int) bool {
	lowBound := rowIndex
	highBound := utility.Min(rowIndex+3, b.C4Rows()-1)
	token := b.C4Grid[rowIndex][columnIndex]
	count := 0

	//no need to go any further if the
	if highBound-lowBound < ConsecutiveTokensForWin-1 {
		return false
	}

	for i := lowBound; i <= highBound; i++ {
		if b.C4Grid[i][columnIndex] == token {
			count++
			if count == ConsecutiveTokensForWin {
				return true
			}
		} else {
			count = 0
		}
	}

	return false
}

//isDownDiagnalWin is a method which determines if the last move was part of a down diagonal win
func (b *ConnectFourBoard) isDownDiagonalWin(rowIndex, columnIndex int) bool {
	// X
	//   X
	//	   X
	//       X
	//set the left-upper bound and right-lower bound
	leftUpperRowIndex, leftUpperColumnIndex := b.leftUpperCoordinate(rowIndex, columnIndex)
	rightLowerRowIndex, rightLowerColumnIndex := b.rightLowerCoordinate(rowIndex, columnIndex)

	//no need to seach further if there are less than 4 tokens to check
	if rightLowerRowIndex-leftUpperRowIndex < ConsecutiveTokensForWin-1 {
		return false
	}

	count := 0
	token := b.C4Grid[rowIndex][columnIndex]
	i, j := leftUpperRowIndex, leftUpperColumnIndex

	for i <= rightLowerRowIndex && j <= rightLowerColumnIndex {
		if b.C4Grid[i][j] == token {
			count++
			if count == ConsecutiveTokensForWin {
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

//isDownDiagnalisUpDiagonalWinWin is a method which determines if the last move was part of a up diagonal win
func (b *ConnectFourBoard) isUpDiagonalWin(rowIndex, columnIndex int) bool {
	//       X
	//	   X
	//   X
	// X
	//set the left-lower and right-upper bounds
	leftLowerRowIndex, leftLowerColumnIndex := b.leftLowerCoordinate(rowIndex, columnIndex)
	rightUpperRowIndex, rightUpperColumnIndex := b.rightUpperCoordinate(rowIndex, columnIndex)

	//test if that there are less than 4 tokens otherwise no need to go further
	if rightUpperColumnIndex-leftLowerColumnIndex < ConsecutiveTokensForWin-1 {
		return false
	}

	count := 0
	token := b.C4Grid[rowIndex][columnIndex]
	i, j := leftLowerRowIndex, leftLowerColumnIndex

	for i >= rightUpperRowIndex && j <= rightUpperColumnIndex {
		if b.C4Grid[i][j] == token {
			count++
			if count == ConsecutiveTokensForWin {
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

//leftLowerCoordinate returns the coordinate of the left-lower cell to evaluate a up diagnal win
func (b *ConnectFourBoard) leftLowerCoordinate(rowIndex, columnIndex int) (int, int) {

	rowUpperBound := b.C4Rows() - 1

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

//rightUpperCoordinate returns the coordinate of the right-upper cell to evaluate a up diagnal win
func (b *ConnectFourBoard) rightUpperCoordinate(rowIndex, columnIndex int) (int, int) {

	columnUpperBound := b.C4Columns() - 1

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
func (b *ConnectFourBoard) leftUpperCoordinate(rowIndex, columnIndex int) (int, int) {
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

//rightLowerCoordinate returns the coordinate of the right-lowercell to evaluate for a down diagnal win
func (b *ConnectFourBoard) rightLowerCoordinate(rowIndex, columnIndex int) (int, int) {

	rowUpperBound := b.C4Rows() - 1
	columnUpperBound := b.C4Columns() - 1

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

func (b *ConnectFourBoard) Grid() [][]enums.Token {
	return b.C4Grid
}

func (b *ConnectFourBoard) LastMove() int {
	return b.C4LastMoveColumnIndex
}

func (b *ConnectFourBoard) Rows() int {
	return len(b.C4Grid)
}

func (b *ConnectFourBoard) Columns() int {
	return len(b.C4Grid[0])
}

