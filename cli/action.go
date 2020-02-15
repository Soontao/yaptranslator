package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Soontao/yaptranslator"
	"github.com/magiconair/properties"
	"github.com/urfave/cli"
)

// Run app
func Run(c *cli.Context) error {
	f := c.GlobalString("file")
	d, e := os.Stat(f)

	if e != nil {
		return e
	}

	authProps := map[string]string{
		"key":      c.GlobalString("key"),
		"secret":   c.GlobalString("secret"),
		"username": c.GlobalString("username"),
		"password": c.GlobalString("password"),
		"region":   c.GlobalString("region"),
		"token":    c.GlobalString("token"),
	}

	providerName := yaptranslator.TranslatorType(c.GlobalString("service"))

	sourceLang := c.GlobalString("lang")

	targetLang := c.GlobalString("target")

	translator, e := yaptranslator.NewTranslator(providerName, authProps)

	if e != nil {
		return e
	}

	if d.IsDir() {
		// list properties files & process files
		return errors.New("Not supported now")
	}

	// process file
	props, e := yaptranslator.TranslatePropertiesFile(
		properties.MustLoadFile(f, properties.UTF8),
		translator,
		sourceLang,
		targetLang,
	)

	if e != nil {
		return e
	}

	nf, e := os.OpenFile(yaptranslator.AddSuffixToFileName(f, fmt.Sprint("_", targetLang)), os.O_RDWR+os.O_CREATE, 0644)

	if e != nil {
		return e
	}

	nf.WriteString(props.String())

	nf.Close()

	return nil
}
