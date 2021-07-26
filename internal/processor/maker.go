package processor

import "os"

func makeDestination(dst string) {
	_ = os.MkdirAll(dst, 0750)
}
