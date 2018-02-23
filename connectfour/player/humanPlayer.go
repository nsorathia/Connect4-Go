package player

import (
	"errors"
	"fmt"
	"games/connectfour/board"
	"games/connectfour/dataDevice"
	"games/connectfour/enums"
	"games/connectfour/utilities"
	"strconv"
)

type HumanPlayer struct {
	id          int
	playerName  string
	playerToken enums.Token
	device      dataDevice.DataDevice
}

func NewHumanPlayer(gameId, id int, name string, token enums.Token, device dataDevice.DataDevice) HumanPlayer {

	return HumanPlayer{
		id:          id,
		playerName:  name,
		playerToken: token,
		device:      device,
	}
}

func (h *HumanPlayer) Name() string {
	return h.playerName
}

func (h *HumanPlayer) Token() enums.Token {
	return h.playerToken
}

func (h *HumanPlayer) Move(board board.Board) (int, error) {

	if board == nil {
		return -1, errors.New("The board object is nil")
	}

	available := board.GetAvailableMoves()
	choice, err := strconv.Atoi("-1")

	for !utility.Contains(available, choice-1) {
		h.device.Write(fmt.Sprintf("%v...Enter a column number for your token", h.playerName))
		input := h.device.Read()
		choice, err = strconv.Atoi(input)
		if err != nil {
			h.device.Write("Try Again...your choice was invalid")
		}
	}
	return choice, nil
}

func (h *HumanPlayer) Id() int {
	return h.id
}
