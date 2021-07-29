package duplicate

import (
	"github.com/nekovalue/unpckr/internal/config"
	"github.com/nekovalue/unpckr/internal/logger"
	"github.com/nekovalue/unpckr/pkg/rename_strategies"
)

// RemoveDuplicates removing all files with same hash received from scanDuplicates
func RemoveDuplicates(config *config.ConfigurationType) {
	duplicatesToRemove := scanDuplicates(config.Storage.Sources)

	indexOffset := 0
	for _, value := range duplicatesToRemove {
		config.Storage.RemoveStoragePairByIndex(value - indexOffset)
		indexOffset++

		logger.Log.Debugf("Duplicates found on %d", value-indexOffset)
	}
	logger.Log.Info("Duplicates removed")
}

// Returns duplicate indexes. It saves one element, that will be original file
func scanDuplicates(paths []string) []int {
	repeats := make(map[string]int)
	var candidatesToRemove []int

	for i, value := range paths {
		hashed, _ := rename_strategies.HashFileMD5(value)
		repeats[hashed]++
		if repeats[hashed] > 1 {
			candidatesToRemove = append(candidatesToRemove, i)
		}
	}

	logger.Log.Info("Duplicates scanned")

	return candidatesToRemove
}
