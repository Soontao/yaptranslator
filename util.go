package yaptranslator

import (
	"bytes"

	"github.com/magiconair/properties"
)

// StringifyProperties with comments
func StringifyProperties(p *properties.Properties) string {
	var w bytes.Buffer
	p.Write(&w, properties.UTF8)
	return w.String()
}
