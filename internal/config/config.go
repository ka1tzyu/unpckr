package config

import (
	"fmt"
	"github.com/nekovalue/unpckr/internal/storage"
)

var CONFIG ConfigurationType

type ConfigurationType struct {
	Storage          storage.Storage
	Sources          *[]string
	Destination      *string
	ConflictRename   *string
	RenameAll        *string
	Pattern          *string
	Log              *bool
	LogLevel         *string
	RemoveDuplicates *bool
	Unzip            *bool
	UnzipPasses      *int
}

func (cfg ConfigurationType) String() string {
	return fmt.Sprintf("ConfigType(Sources: %v; Destination: %v; ConflictRename: %v; RenameAll: %v; "+
		"Pattern: %v; Log: %v; LogLevel: %v; RemoveDuplicates: %v; Unzip: %v; UnzipPasses: %v)",
		*cfg.Sources, *cfg.Destination, *cfg.ConflictRename, *cfg.RenameAll, *cfg.Pattern, *cfg.Log, *cfg.LogLevel,
		*cfg.RemoveDuplicates, *cfg.Unzip, *cfg.UnzipPasses)
}
