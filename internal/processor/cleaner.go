package processor

import "os"

func cleaner() {
	_ = os.RemoveAll("__tmp__")
}
