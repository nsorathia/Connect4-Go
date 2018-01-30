package player

import (
	"errors"
	"fmt"
	"games/abstract"
	"games/connectfour/enums"
	"games/connectfour/utilities"
	"strconv"
)

type HumanPlayer struct {
	playerName  string
	playerToken enums.Token
	device      abstract.DataDevice
}

func NewHumanPlayer(n string, t enums.Token, d abstract.DataDevice) HumanPlayer {
	return HumanPlayer{
		playerName:  n,
		playerToken: t,
		device:      d,
	}
}

func (h *HumanPlayer) Name() string {
	return h.playerName
}

func (h *HumanPlayer) Token() enums.Token {
	return h.playerToken
}

func (h *HumanPlayer) Move(board abstract.Board) (int, error) {

	if board == nil {
		return -1, errors.New("The board object is nil")
	}

	available := board.GetAvailableMoves()
	choice, err := strconv.Atoi("-1")

	for !utility.Contains(available, choice) {
		h.device.Write(fmt.Sprintf("%v...Enter a column number for your token", h.playerName))
		input := h.device.Read()
		choice, err = strconv.Atoi(input)
		if err != nil {
			h.device.Write("Try Again...your choice was invalid")
		}
	}
	return choice, nil
}
