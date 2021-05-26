package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ka1i/wispeeer/internal/app/cmd"
	"github.com/ka1i/wispeeer/internal/pkg/usage"
	"github.com/ka1i/wispeeer/internal/pkg/utils"
	"github.com/ka1i/wispeeer/pkg/version"
)

var (
	success = 0
	failure = 137
)

func main() {
	os.Exit(wispeeer())
}

func wispeeer() int {
	if len(os.Args) > 1 {
		var argc = len(os.Args)
		var argv = os.Args[1:]
		start(argc, argv)
	} else {
		usage.Usage()
		return failure
	}
	return success
}

func start(argc int, argv []string) {
	var err error

	switch argv[0] {
	case "-i", "init":
		if argc > 2 {
			if utils.IsValid(argv[1]) {
				err = cmd.Initialzation(argv[1])
			} else {
				err = fmt.Errorf("invalid name")
			}
		} else {
			err = fmt.Errorf("wispeeer init <ka1i.github.io>")
		}

	case "-n", "new":
		if argc > 2 {
			if utils.IsValid(argv[1]) {
				log.Println("new")
			} else {
				err = fmt.Errorf("invalid title")
			}
		} else {
			err = fmt.Errorf("wispeeer new <title>")
		}
	case "-g", "generate":
		log.Println("generate")
	case "-s", "server":
		log.Println("server")
	case "-d", "deploy":
		log.Println("deploy")
	case "-h", "--help", "help":
		usage.Usage()
	case "-v", "--version", "version":
		version.Version()
	default:
		err = fmt.Errorf("wispeeer usage: wispeeer -h")
	}
	if err != nil {
		log.Printf("%s", err)
	}
}
