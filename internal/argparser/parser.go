package argparser

import (
	"github.com/akamensky/argparse"
	"github.com/nekovalue/unpckr/internal/config"
)

// ParseArguments is used to process arguments and pass data to config
func ParseArguments(args []string, cfg *config.ConfigurationType) (argparse.Parser, error) {
	parser := argparse.NewParser("unpckr-cli", "Provides some unpacking functions")

	cfg.Sources = parser.StringList("s", "source",
		&argparse.Options{
			Required: true,
			Validate: sourceFoldersValidate,
			Help:     "Source folder, at least one",
		})
	cfg.Destination = parser.String("d", "dest",
		&argparse.Options{
			Required: false,
			Validate: destinationFolderValidate,
			Help:     "Destination folder",
			Default:  "temp",
		})
	cfg.ConflictRename = parser.Selector("c", "conflict-rename", []string{"simpleRandom", "none"},
		&argparse.Options{
			Required: false,
			Help:     "Conflict rename strategy :[simpleRandom - add to filename 10 random characters]",
			Default:  "simpleRandom",
		})
	cfg.RenameAll = parser.Selector("r", "rename-all", []string{"hash", "none"},
		&argparse.Options{
			Required: false,
			Help:     "Rename all files through selected strategy :[hash - file get his hash string as name]",
			Default:  "none",
		})
	cfg.Pattern = parser.String("p", "pattern",
		&argparse.Options{
			Required: false,
			Help:     "Rename all files trough specified pattern",
			Default:  "none",
		})
	cfg.Log = parser.Flag("l", "log",
		&argparse.Options{
			Required: false,
			Help:     "Show logging stream",
			Default:  false,
		})
	cfg.LogLevel = parser.Selector("", "log-level", []string{"DEBUG", "INFO", "WARN", "ERROR", "none"},
		&argparse.Options{
			Required: false,
			Help:     "Logging level",
			Default:  "none",
		})
	cfg.RemoveDuplicates = parser.Flag("", "remove-duplicates",
		&argparse.Options{
			Required: false,
			Help:     "Remove duplicates basing on file`s hash",
			Default:  false,
		})
	cfg.Unzip = parser.Flag("z", "unzip",
		&argparse.Options{
			Required: false,
			Help:     "Unzip all archives from the sources",
			Default:  false,
		})
	cfg.UnzipPasses = parser.Int("", "unzip-passes",
		&argparse.Options{
			Required: false,
			Help:     "How many times pass the source files for unzip",
			Default:  1,
		})
	cfg.RemoveSources = parser.Flag("", "remove-sources",
		&argparse.Options{
			Required: false,
			Help:     "If enabled delete all sources after work",
			Default:  false,
		})

	err := parser.Parse(args)

	return *parser, err
}
