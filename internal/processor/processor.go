package processor

import (
	"fmt"
	"github.com/nekovalue/unpckr/internal/config"
	"github.com/nekovalue/unpckr/internal/copy"
	"github.com/nekovalue/unpckr/internal/rename"
)

func Process(config *config.ConfigurationType) error {
	err := scanSources(config)
	if err != nil {
		return err
	}

	err = generateDestinations(config)
	if err != nil {
		return err
	}

	switch {
	case *config.ConflictRename == "simpleRandom":
		rename.RandomizeConflicts(config)
	}

	err = copy.WorkAll(config)
	if err != nil {
		return err
	}

	fmt.Println("Done!")

	return nil
}
