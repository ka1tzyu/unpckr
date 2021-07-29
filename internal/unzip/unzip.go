package unzip

import (
	"archive/zip"
	"fmt"
	"github.com/nekovalue/unpckr/internal/config"
	"github.com/nekovalue/unpckr/internal/logger"
	"github.com/nekovalue/unpckr/internal/scanner"
	path2 "github.com/nekovalue/unpckr/pkg/path"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// SourcesUnzip passes file unzipping, removes zip files from source, scan new sources
func SourcesUnzip(config *config.ConfigurationType) {
	for i := 0; i < *config.UnzipPasses; i++ {
		zips := scanZips(config.Storage.Sources)
		unzipped := filesUnzip(zips)
		for zipPath, dest := range unzipped {
			config.Storage.RemoveStoragePairByIndex(config.Storage.GetIndexOfSource(zipPath))
			_ = scanner.ScanDirectory(dest, config)
			config.Storage.ClearStorageDestinations()
			_ = config.Storage.GenerateDestinations(*config.Destination)
		}
	}

	logger.Log.Info("Sources were unzipped")
}

// Unzip all zips and returns map[zipPath]outputFolder
func filesUnzip(zips []string) map[string]string {
	result := make(map[string]string)

	for _, value := range zips {
		result[value], _ = fileUnzip(value)
	}

	return result
}

// Unzip single archive, returns (outputPath, error)
func fileUnzip(path string) (string, error) {
	dest := "__tmp__" + string(os.PathSeparator) + path2.GetFileNameWithoutExtensionFromPath(path)

	r, err := zip.OpenReader(path)
	if err != nil {
		return dest, err
	}

	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	err = os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}

		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			err = os.MkdirAll(path, f.Mode())
		} else {
			err = os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}

			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return dest, err
		}
	}

	return dest, nil
}

// Returns all zips inside paths
func scanZips(paths []string) []string {
	var zips []string

	for _, value := range paths {
		if value[len(value)-4:] == ".zip" {
			zips = append(zips, value)
			logger.Log.Debugf("New zip was found {%s}", value)
		}
	}

	logger.Log.Info("Archives were unzipped")

	return zips
}
