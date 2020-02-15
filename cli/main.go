package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

// Version string, in release version
// This variable will be overwrited by complier
var Version = "SNAPSHOT"

// AppName of this application
var AppName = "Yet Another Properties (file) Translator"

// AppUsage of this application
var AppUsage = "A Command Line Tool"

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Name = AppName
	app.Usage = AppUsage
	app.Flags = options
	app.Action = Run

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
