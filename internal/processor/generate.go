package processor

import (
	"github.com/nekovalue/unpckr/internal/config"
	"os"
	"strings"
)

func generateDestinations(config *config.ConfigurationType) error {
	for _, src := range config.Storage.Sources {
		filename := strings.Split(src, string(os.PathSeparator))
		dest := *config.Destination + string(os.PathSeparator) + filename[len(filename)-1]
		config.Storage.AppendDestination(dest)
	}
	return nil
}
