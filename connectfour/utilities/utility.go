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


//Min is a helper function which returns the minimun of two integers
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//Max is a helper function which returns the maximum of two integers
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
