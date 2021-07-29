package app

import (
	"fmt"
	"github.com/nekovalue/unpckr/internal/config"
	"log"
	"os"

	"github.com/nekovalue/unpckr/internal/argparser"
	"github.com/nekovalue/unpckr/internal/logger"
	"github.com/nekovalue/unpckr/internal/processor"
)

var CONFIG config.ConfigurationType

// Run App entry point
func Run(args []string) {
	// Parse arguments and put data to config
	parser, err := argparser.ParseArguments(args, &CONFIG)
	if err != nil && &parser != nil {
		fmt.Println(parser.Usage(err))
		os.Exit(0)
	}

	// Global logging setting
	logger.SetLogger(logger.Log, *CONFIG.Log, *CONFIG.LogLevel)

	// Process filters
	err = processor.Process(&CONFIG)
	if err != nil {
		log.Fatal(err)
	}
}
