package processor

import (
	"os"

	"github.com/nekovalue/unpckr/internal/config"
)

func cleanTemp() {
	_ = os.RemoveAll("__tmp__")
}

func cleanSources(config config.ConfigurationType) {
	for _, value := range *config.Sources {
		_ = os.RemoveAll(value)
	}
}
