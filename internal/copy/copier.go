package copy

import (
	"io"
	"os"

	"github.com/nekovalue/unpckr/internal/config"
	"github.com/nekovalue/unpckr/internal/logger"
)

func WorkAll(config *config.ConfigurationType) error {
	for i, src := range config.Storage.Sources {
		err := singleCopy(src, config.Storage.Destinations[i])
		if err != nil {
			return err
		}
	}

	logger.Log.Info("All files was copied")

	return nil
}

func singleCopy(src, dst string) error {
	originalFile, err := os.Open(src)
	if err != nil {
		return err
	}

	defer func(originalFile *os.File) {
		err = originalFile.Close()
		if err != nil {
			logger.Log.Fatal(err)
		}
	}(originalFile)

	newFile, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer func(newFile *os.File) {
		err = newFile.Close()
		if err != nil {
			logger.Log.Fatal(err)
		}
	}(newFile)

	_, err = io.Copy(newFile, originalFile)
	if err != nil {
		return err
	}

	logger.Log.Debugf("{%s} was copied to {%s}", src, dst)

	return nil
}
