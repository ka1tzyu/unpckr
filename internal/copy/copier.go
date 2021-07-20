package copy

import (
	"github.com/nekovalue/unpckr/internal/config"
	"io"
	"os"
)

func WorkAll(config *config.ConfigurationType) error {
	for i, src := range config.Storage.Sources {
		err := singleCopy(src, config.Storage.Destinations[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func singleCopy(src, dst string) error {
	originalFile, err := os.Open(src)
	if err != nil {
		return err
	}

	defer func(originalFile *os.File) {
		_ = originalFile.Close()
	}(originalFile)

	newFile, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer func(newFile *os.File) {
		_ = newFile.Close()
	}(newFile)

	_, err = io.Copy(newFile, originalFile)
	if err != nil {
		return err
	}

	return nil
}
