package dataDevice

import (
	"bufio"
	"fmt"
	"games/abstract"
	"os"
)

func init() {
	abstract.NewDataDevice = NewIODevice
}

func NewIODevice() abstract.DataDevice {
	return &IODevice{}
}

type IODevice struct {
}

func (i *IODevice) Write(input string) {
	fmt.Println(input)
}

func (i *IODevice) Read() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}
