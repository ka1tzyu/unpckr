package app

import (
	"fmt"
	"github.com/nekovalue/unpckr/internal/argparser"
	"github.com/nekovalue/unpckr/internal/config"
	"github.com/nekovalue/unpckr/internal/processor"
	"os"
)

func Run(args []string) {
	parser, err := argparser.ParseArguments(args, &config.CONFIG)
	if err != nil && &parser != nil {
		fmt.Println(parser.Usage(err))
		os.Exit(0)
	}

	err = processor.Process(&config.CONFIG)
	if err != nil {
		fmt.Println("appError: ", err)
	}
}
