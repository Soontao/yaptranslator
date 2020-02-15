package yaptranslator

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/magiconair/properties"
	"gopkg.in/cheggaaa/pb.v2"
)

// StringifyProperties with comments
func StringifyProperties(p *properties.Properties) string {
	var w bytes.Buffer
	p.Write(&w, properties.UTF8)
	return w.String()
}

// TranslatePropertiesFile process
func TranslatePropertiesFile(props *properties.Properties, translator Translator, lang, target string) (*properties.Properties, error) {
	rt := properties.NewProperties()

	propsMap := props.Map()

	progress := pb.StartNew(len(propsMap))

	for k, v := range propsMap {
		tv, err := translator.GetTranslation(v, lang, target)
		if err == nil {
			rt.SetValue(k, tv)
		} else {
			log.Printf("translate '%s' failed from %s to %s: %v", v, lang, target, err)
		}
		progress.Increment()
	}

	progress.Finish()

	return rt, nil
}

// AddSuffixToFileName string
func AddSuffixToFileName(path, suffix string) string {
	base, filename := filepath.Split(path)
	ext := filepath.Ext(filename)
	filebasename := strings.TrimRight(filename, ext)
	return fmt.Sprint(base, filebasename, suffix, ext)
}
