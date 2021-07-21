package duplicate

import (
	"github.com/nekovalue/unpckr/internal/config"
	"github.com/nekovalue/unpckr/internal/rename"
)

func RemoveDuplicates(config *config.ConfigurationType) {
	duplicatesToRemove := scanDuplicates(config.Storage.Sources)

	indexOffset := 0
	for _, value := range duplicatesToRemove {
		config.Storage.RemoveStoragePairByIndex(value - indexOffset)
		indexOffset++
	}
}

// Returns duplicate indexes. It saves one element, that will be original file
func scanDuplicates(paths []string) []int {
	repeats := make(map[string]int)
	var candidatesToRemove []int

	for i, value := range paths {
		hashed, _ := rename.HashFileMD5(value)
		repeats[hashed]++
		if repeats[hashed] > 1 {
			candidatesToRemove = append(candidatesToRemove, i)
		}
	}

	return candidatesToRemove
}
