package main

import (
	"log"
)

var makeCmd = &Command{
	Execute:     make,
	Name:        "make",
	Usage:       "alkalisk make [type] [arguments]",
	Description: "Create something new!",
}

func make(args []string) {

	if len(args) < 2 {
		usage()
	}

	if args[1] == "model" {
		log.Println("model")
	}
}
