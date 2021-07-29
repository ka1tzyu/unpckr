package processor

import "os"

// Creates destination folder
func makeDestination(dst string) {
	_ = os.MkdirAll(dst, 0750)
}
