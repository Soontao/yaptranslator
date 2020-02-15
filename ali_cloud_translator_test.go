package yaptranslator

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetAliCloudTranslatorFromEnv() (Translator, error) {
	// create translator
	keyID := os.Getenv("TEST_ALICLOUD_KEY_ID")
	keySecret := os.Getenv("TEST_ALICLOUD_KEY_SECRET")
	region := os.Getenv("TEST_ALICLOUD_KEY_REGION")

	return NewTranslator(ALICLOUD, map[string]string{
		"key":    keyID,
		"secret": keySecret,
		"region": region,
	})

}

func TestNewAliCloudTranslator(t *testing.T) {

	// export TEST_ALICLOUD_KEY_ID=
	// export TEST_ALICLOUD_KEY_SECRET=
	// export TEST_ALICLOUD_KEY_REGION=cn-hangzhou

	a := assert.New(t)
	c, e := GetAliCloudTranslatorFromEnv()

	a.NoError(e)

	tr, e := c.GetTranslation("你好", "zh", "en")
	a.NoError(e)
	a.Equal("Hello", tr)

	tr, e = c.GetTranslation("Hello", "en", "zh")
	a.NoError(e)
	a.Equal("你好", tr)

	_, e = c.GetTranslation("你不好", "not exist", "cn")
	a.Error(e) // throw error
	a.Equal("translate from source to target not support", e.Error())
}
