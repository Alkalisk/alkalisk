package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var newCmd = &Command{
	Execute:     new,
	Name:        "new",
	Usage:       "alkalisk new projectname",
	Description: "Creates a new project in your working directory.",
}

func new(args []string) {

	if len(args) < 2 {
		usage()
	}

	getDependencies()
	createFiles(args[1])

	fmt.Println("Application succesfully created")
}

func createFiles(name string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dest := filepath.Join(dir, name)

	src := filepath.Join(os.Getenv("GOPATH"), "src/github.com/Alkalisk/alkalisk/skeleton")

	err = CopyDir(src, dest)
	if err != nil {
		log.Fatal(err)
	}
}

func getDependencies() {
	err, _ := exec.Command("go", "get", "github.com/Alkalisk/alkalisk").Output()
	if err != nil {
		//log.Panic(err)
	}
}
