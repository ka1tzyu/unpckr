package tools

// RemoveSliceElementByIndex returns new slice without element of given index
func RemoveSliceElementByIndex(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

// EqualSlice checks slices to equality
func EqualSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
