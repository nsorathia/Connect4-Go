package dataDevice

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

func init() {
	NewDataDevice = NewIODevice
}

func NewIODevice() DataDevice {
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
	return strings.Replace(text, "\n", "", -1)
}
