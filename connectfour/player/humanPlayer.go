package player

import (
	"games/connectfour/dataDevice"
	"errors"
	"fmt"
	"games/connectfour/board"
	"games/connectfour/enums"
	"games/connectfour/utilities"
	"strconv"
)

type HumanPlayer struct {
	playerName  string
	playerToken enums.Token
	device      dataDevice.DataDevice
}

func NewHumanPlayer(name string, token enums.Token, device dataDevice.DataDevice) HumanPlayer {
	return HumanPlayer{
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
