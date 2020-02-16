package yaptranslator

import (
	"testing"

	"github.com/magiconair/properties"
	"github.com/stretchr/testify/assert"
)

func TestAddSuffixToFileName(t *testing.T) {
	a := assert.New(t)
	a.Equal("./a/b/c/d_zh.properties", AddSuffixToFileName("./a/b/c/d.properties", "_zh"))
	a.Equal("./a/b/c/descriptions_zh.properties", AddSuffixToFileName("./a/b/c/descriptions.properties", "_zh"))
}

func TestTranslatePropertiesFile(t *testing.T) {
	a := assert.New(t)

	tr, e := GetAliCloudTranslatorFromEnv()

	a.NoError(e)
	p := properties.MustLoadFile("./resources/example.2.properties", properties.UTF8)

	tresult, e := TranslatePropertiesFile(p, tr, "en", "zh")
	a.NoError(e)

	a.Equal("你好世界", tresult.MustGetString("a.b.c.d"))
}
