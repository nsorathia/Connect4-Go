package dataDevice

//DataDevice is an abstract representation of some IO device to interact with the Players
type DataDevice interface {
	Write(input string)
	Read() string
}

type deviceFactory func() DataDevice

var NewDataDevice deviceFactory
