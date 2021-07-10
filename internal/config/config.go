package config

import (
	"fmt"
)

var CONFIG configType

type configType struct {
	Sources          *[]string
	Destination      *string
	ConflictRename   *string
	RenameAll        *string
	Pattern          *string
	RemoveDuplicates *bool
	Unzip            *bool
}

func (cfg configType) String() string {
	return fmt.Sprintf("ConfigType(Sources: %v; Destination: %v; ConflictRename: %v; RenameAll: %v; "+
		"Pattern: %v; RemoveDuplicates: %v; Unzip: %v)", *cfg.Sources, *cfg.Destination, *cfg.ConflictRename,
		*cfg.RenameAll, *cfg.Pattern, *cfg.RemoveDuplicates, *cfg.Unzip)
}
