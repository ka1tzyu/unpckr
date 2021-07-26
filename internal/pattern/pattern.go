package pattern

import (
	"github.com/nekovalue/unpckr/internal/logger"
	"github.com/nekovalue/unpckr/pkg/path"
	"strconv"
	"strings"

	"github.com/nekovalue/unpckr/internal/config"
)

func RenameByPattern(pattern string, config *config.ConfigurationType) {
	if pattern == "numeric" {
		for i := range config.Storage.Destinations {
			extParts := strings.Split(config.Storage.Destinations[i], ".")
			ext := "." + extParts[len(extParts)-1]
			config.Storage.Destinations[i] =
				path.GetPathWithoutFileName(config.Storage.Destinations[i]) + strconv.Itoa(i) + ext

			logger.Log.Debugf("New name {%s} was assigned", config.Storage.Destinations[i])
		}
	} else {
		if strings.Contains(pattern, "%n") {
			for i := range config.Storage.Destinations {
				extParts := strings.Split(config.Storage.Destinations[i], ".")
				ext := "." + extParts[len(extParts)-1]
				name := strings.Replace(pattern, "%n", strconv.Itoa(i), 1)
				config.Storage.Destinations[i] =
					path.GetPathWithoutFileName(config.Storage.Destinations[i]) + name + ext
				logger.Log.Debugf("New name {%s} was assigned", config.Storage.Destinations[i])
			}
		}
	}

	logger.Log.Infof("Files was renames by pattern {%s}", pattern)
}
