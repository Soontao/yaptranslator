package main

import (
	"os"

	"github.com/urfave/cli"
)

// Run app
func Run(c *cli.Context) error {
	f := c.GlobalString("file")
	d, e := os.Stat(f)
	if e != nil {
		return e
	}
	if d.IsDir() {
		// list properties files & process files
	} else {
		// process file
	}
	return nil
}
