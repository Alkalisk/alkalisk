package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var runCmd = &Command{
	Execute:     run,
	Name:        "run",
	Usage:       "TODO",
	Description: "TODO",
}

func run(args []string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	mainDest := filepath.Join(dir, "main.go")

	f, err := os.Create(mainDest)

	check(err)

	defer f.Close()

	w := bufio.NewWriter(f)

	_, err = fmt.Fprintf(w, mainTmpl)

	check(err)

}

const mainTmpl string = `
package main

func main() {
	registerRouters()
}
`
