package processor

import (
	"fmt"
	"github.com/nekovalue/unpckr/internal/config"
	"github.com/nekovalue/unpckr/internal/copy"
	"github.com/nekovalue/unpckr/internal/scanner"
)

func Process(config *config.ConfigurationType) error {
	err := scanner.ScanSources(config)
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

	cleaner()

	fmt.Println("All work done!")

	return nil
}
