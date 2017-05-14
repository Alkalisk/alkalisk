package main

import (
	"flag"
	"log"
	"os"
)

type Command struct {
	Execute     func(args []string)
	Name        string
	Usage       string
	Description string
}

var commands = []*Command{
	newCmd,
	makeCmd,
	installCmd,
	runCmd,
}

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		usage()
	}

	for _, cmd := range commands {
		if cmd.Name == flag.Args()[0] {
			cmd.Execute(flag.Args())
		}
	}
}

func usage() {
	log.Println(usageTemplate)
	os.Exit(0)
}

const usageTemplate string = `
	usage: alkalisk command [arguments]
`
