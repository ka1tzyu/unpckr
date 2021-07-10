package argparser

import (
	"github.com/akamensky/argparse"
	"github.com/nekovalue/unpckr/internal/config"
)

func ParseArguments(args []string) (argparse.Parser, error) {
	parser := argparse.NewParser("unpckr-cli", "Provides some unpacking functions")

	config.CONFIG.Sources = parser.StringList("s", "source",
		&argparse.Options{
			Required: true,
			Validate: sourceFoldersValidate,
			Help:     "Source folder, at least one",
		})
	config.CONFIG.Destination = parser.String("d", "dest",
		&argparse.Options{
			Required: false,
			Validate: destinationFolderValidate,
			Help:     "Destination folder",
			Default:  "temp/",
		})
	config.CONFIG.ConflictRename = parser.Selector("c", "conflict-rename", []string{"simpleRandom"},
		&argparse.Options{
			Required: false,
			Help:     "Conflict rename strategy :[simpleRandom - add to filename 10 random characters]",
			Default:  "simpleRandom",
		})
	config.CONFIG.RenameAll = parser.Selector("r", "rename-all", []string{"hash", "none"},
		&argparse.Options{
			Required: false,
			Help:     "Rename all files through selected strategy :[hash - file get his hash string as name]",
			Default:  "none",
		})
	config.CONFIG.Pattern = parser.String("p", "pattern",
		&argparse.Options{
			Required: false,
			Validate: patternValidate,
			Help:     "Rename all files trough specified pattern",
			Default:  "none",
		})
	config.CONFIG.RemoveDuplicates = parser.Flag("", "remove-duplicates",
		&argparse.Options{
			Required: false,
			Help:     "Remove duplicates basing on file`s hash",
			Default:  false,
		})
	config.CONFIG.Unzip = parser.Flag("z", "unzip",
		&argparse.Options{
			Required: false,
			Help:     "Unzip all archives from the sources",
			Default:  false,
		})

	err := parser.Parse(args)

	return *parser, err
}
