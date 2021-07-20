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
			newFileNameParts := strings.Split(config.Storage.Destinations[i], ".")
			newFileName := newFileNameParts[0] + "__" +
				generateRandomString(10, i) + "." + newFileNameParts[len(newFileNameParts)-1]
			config.Storage.Destinations[i] = newFileName
		}
	}
}
