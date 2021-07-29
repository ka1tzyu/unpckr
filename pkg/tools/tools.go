package tools

// RemoveSliceElementByIndex returns new slice without element of given index
func RemoveSliceElementByIndex(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
