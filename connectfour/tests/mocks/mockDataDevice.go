package mocks

import (
	"games/connectfour/dataDevice"
	"github.com/stretchr/testify/mock"
)

func NewIODevice() dataDevice.DataDevice {
	return &MockDataDevice{}
}

type MockDataDevice struct {
	mock.Mock
}

func (i *MockDataDevice) Write(input string) {
	i.Called()
}

func (i *MockDataDevice) Read() string {
	args := i.Called()
	return args.String(0)
}