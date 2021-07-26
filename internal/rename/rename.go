package rename

import (
	"github.com/nekovalue/unpckr/internal/config"
	"github.com/nekovalue/unpckr/internal/logger"
	"github.com/nekovalue/unpckr/pkg/path"
	"github.com/nekovalue/unpckr/pkg/random"
	"github.com/nekovalue/unpckr/pkg/rename_strategies"
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
				random.GenerateRandomString(10, i) + "." + newFileNameParts[len(newFileNameParts)-1]
			config.Storage.Destinations[i] = newFileName

			logger.Log.Debugf("New name {%s} was assigned", newFileName)
		}
	}

	logger.Log.Info("Conflicts were randomized")
}

func HashingDestinations(config *config.ConfigurationType) error {
	for i, value := range config.Storage.Destinations {
		hash, _ := rename_strategies.HashFileMD5(config.Storage.Sources[i])

		newFileNameParts := strings.Split(value, ".")
		newFileName := path.GetPathWithoutFileName(value) + hash + "." + newFileNameParts[len(newFileNameParts)-1]

		config.Storage.Destinations[i] = newFileName

		logger.Log.Debugf("New name {%s} was assigned", newFileName)
	}

	logger.Log.Info("Destination names were hashed")

	return nil
}
