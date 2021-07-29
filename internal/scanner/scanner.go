package scanner

import (
	"github.com/nekovalue/unpckr/internal/config"
	"github.com/nekovalue/unpckr/internal/logger"
	"io/fs"
	"path/filepath"
)

// ScanSources scans all sources from [config.Sources]
func ScanSources(config *config.ConfigurationType) error {
	for _, src := range *config.Sources {
		err := ScanDirectory(src, config)
		if err != nil {
			return err
		}
	}

	logger.Log.Info("Sources were scanned")

	return nil
}

// ScanDirectory walks through source folders and append files to [config.Storage.Sources]
func ScanDirectory(dir string, config *config.ConfigurationType) error {
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err == nil {
			if !d.IsDir() {
				config.Storage.AppendSource(path)
			}
		}
		return err
	})
	return err
}
