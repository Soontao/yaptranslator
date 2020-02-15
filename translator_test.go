package yaptranslator

import (
	"testing"

	"github.com/magiconair/properties"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestPropertyString(t *testing.T) {
	a := assert.New(t)
	p := properties.MustLoadFile("./resources/example.1.properties", properties.UTF8)
	a.Equal("value1", p.GetString("key1", ""))
	a.Equal("value2", p.GetString("key2", ""))
	a.Equal("3", p.GetString("root.sub.sub2", ""))
	p.MustSet("root.sub.sub2", "4")
	a.Equal("key1 = value1\nkey2 = value2\nroot.sub.sub2 = 4\n", StringifyProperties(p))
}

func TestLangugeCode(t *testing.T) {
	a := assert.New(t)
	a.Equal("zh-Hans", language.SimplifiedChinese.String())
	a.Equal("zh", language.Chinese.String())
	a.Equal("en-US", language.AmericanEnglish.String())
	a.Equal("en-GB", language.BritishEnglish.String())
}
