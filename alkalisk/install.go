package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var installCmd = &Command{
	Execute:     install,
	Name:        "install",
	Usage:       "TODO",
	Description: "TODO",
}

func install(args []string) {
	if len(args) < 2 {
		usage()
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	switch args[1] {
	case "router":
		err, _ := exec.Command("go", "get -d -u", "github.com/Alkalisk/router").Output()
		if err != nil {
			//	log.Panic(err)
		}

		_, createErr := os.Create(dir + "/routes/web.go")
		if createErr != nil {
			log.Fatalln("Something went wrong, please make sure Alkalisk has been properly initialized in this directory.")
			panic(createErr)
		}

		log.Println("Succesfully installed Alkalisk/router")
		break
	}
}
