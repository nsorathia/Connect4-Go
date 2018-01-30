package utility

//Contains is a function that chekcs if a given integer value is found within a slice
func Contains(slice []int, element int) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == element {
			return true
		}
	}
	return false
}
