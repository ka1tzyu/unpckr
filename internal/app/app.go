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

func Run(args []string) {
	parser, err := argparser.ParseArguments(args, &CONFIG)
	if err != nil && &parser != nil {
		fmt.Println(parser.Usage(err))
		os.Exit(0)
	}

	logger.SetLogger(logger.Log, *CONFIG.Log, *CONFIG.LogLevel)

	err = processor.Process(&CONFIG)
	if err != nil {
		log.Fatal(err)
	}
}
