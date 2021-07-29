package argparser

import (
	"fmt"
	"os"
)

// Check if the folders exist
func sourceFoldersValidate(args []string) error {
	if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		return fmt.Errorf("validate: source %v is not exist", args[0])
	}
	return nil
}

// Check if the folder exists, create if it`s not
func destinationFolderValidate(args []string) error {
	if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		err := os.MkdirAll(args[0], 0750)
		return err
	}
	return nil
}
