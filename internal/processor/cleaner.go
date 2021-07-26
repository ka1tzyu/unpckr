package processor

import (
	"github.com/nekovalue/unpckr/internal/logger"
	"os"

	"github.com/nekovalue/unpckr/internal/config"
)

func cleanTemp() {
	_ = os.RemoveAll("__tmp__")

	logger.Log.Debug("Temp folder was cleaned")
}

func cleanSources(config config.ConfigurationType) {
	for _, value := range *config.Sources {
		_ = os.RemoveAll(value)

	}
	logger.Log.Debug("Sources were cleaned")
}
