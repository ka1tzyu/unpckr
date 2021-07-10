package app

import (
	"fmt"
	"github.com/nekovalue/unpckr/internal/argparser"
	"github.com/nekovalue/unpckr/internal/config"
	"os"
)

func Run(args []string) {
	parser, err := argparser.ParseArguments(args)
	if err != nil && &parser != nil {
		fmt.Println(parser.Usage(err))
		os.Exit(0)
	}
	fmt.Println(config.CONFIG)
}
