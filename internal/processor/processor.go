package processor

import (
	"fmt"
	"github.com/nekovalue/unpckr/internal/config"
	"github.com/nekovalue/unpckr/internal/copy"
)

func Process(config *config.ConfigurationType) error {
	err := scanSources(config)
	if err != nil {
		return err
	}

	err = config.Storage.GenerateDestinations(*config.Destination)
	if err != nil {
		return err
	}

	err = checkFilters(config)
	if err != nil {
		return err
	}

	err = copy.WorkAll(config)
	if err != nil {
		return err
	}

	fmt.Println("All work done!")

	return nil
}
