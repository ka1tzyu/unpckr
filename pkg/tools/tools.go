package tools

func RemoveSliceElementByIndex(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
