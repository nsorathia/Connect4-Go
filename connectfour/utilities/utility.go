package utility

import (
	"math/rand"
	"time"
)


//Contains is a function that chekcs if a given integer value is found within a slice
func Contains(slice []int, element int) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == element {
			return true
		}
	}
	return false
}

//Random returns a random integer between the given min and max
func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
