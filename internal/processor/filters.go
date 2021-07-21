package processor

import (
	"github.com/nekovalue/unpckr/internal/config"
	"github.com/nekovalue/unpckr/internal/duplicate"
	"github.com/nekovalue/unpckr/internal/rename"
	"github.com/nekovalue/unpckr/internal/unzip"
)

func checkFilters(config *config.ConfigurationType) error {
	if *config.Unzip {
		unzip.SourcesUnzip(config)
	}

	if *config.RemoveDuplicates {
		duplicate.RemoveDuplicates(config)
	}

	if *config.RenameAll == "hash" {
		err := rename.HashingDestinations(config)
		if err != nil {
			return err
		}
	}

	if *config.ConflictRename == "simpleRandom" {
		rename.RandomizeConflicts(config)
	}

	return nil
}
