package rename

import (
	"github.com/nekovalue/unpckr/internal/config"
	"strings"
)

func RandomizeConflicts(config *config.ConfigurationType) {
	repeats := make(map[string]int)

	for _, value := range config.Storage.Destinations {
		repeats[value]++
	}

	for i, value := range config.Storage.Destinations {
		if repeats[value] > 1 {
			newFileNameParts := strings.Split(value, ".")
			newFileName := strings.Join(newFileNameParts[0:len(newFileNameParts)-1], ".") + "__" +
				generateRandomString(10, i) + "." + newFileNameParts[len(newFileNameParts)-1]
			config.Storage.Destinations[i] = newFileName
		}
	}
}

func HashingDestinations(config *config.ConfigurationType) error {
	for i, value := range config.Storage.Destinations {
		hash, _ := hashFileMD5(config.Storage.Sources[i])

		newFileNameParts := strings.Split(value, ".")
		newFileName := getPathWithoutFileName(value) + hash + "." + newFileNameParts[len(newFileNameParts)-1]

		config.Storage.Destinations[i] = newFileName
	}

	return nil
}
