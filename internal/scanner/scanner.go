package scanner

import (
	"github.com/nekovalue/unpckr/internal/config"
	"io/fs"
	"path/filepath"
)

func ScanSources(config *config.ConfigurationType) error {
	for _, src := range *config.Sources {
		err := ScanDirectory(src, config)
		if err != nil {
			return err
		}
	}
	return nil
}

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
