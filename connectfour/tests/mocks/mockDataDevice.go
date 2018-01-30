package mocks

import (
	"github.com/stretchr/testify/mock"
	"games/abstract"
)

func NewIODevice() abstract.DataDevice {
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