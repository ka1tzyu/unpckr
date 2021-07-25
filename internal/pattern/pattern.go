package pattern

import (
	"strconv"
	"strings"

	"github.com/nekovalue/unpckr/internal/config"
	"github.com/nekovalue/unpckr/internal/rename"
)

func RenameByPattern(pattern string, config *config.ConfigurationType) {
	if pattern == "numeric" {
		for i := range config.Storage.Destinations {
			extParts := strings.Split(config.Storage.Destinations[i], ".")
			ext := "." + extParts[len(extParts)-1]
			config.Storage.Destinations[i] = rename.GetPathWithoutFileName(config.Storage.Destinations[i]) + strconv.Itoa(i) + ext
		}
	} else {
		if strings.Contains(pattern, "%n") {
			for i := range config.Storage.Destinations {
				extParts := strings.Split(config.Storage.Destinations[i], ".")
				ext := "." + extParts[len(extParts)-1]
				name := strings.Replace(pattern, "%n", strconv.Itoa(i), 1)
				config.Storage.Destinations[i] = rename.GetPathWithoutFileName(config.Storage.Destinations[i]) + name + ext
			}
		}
	}
}
