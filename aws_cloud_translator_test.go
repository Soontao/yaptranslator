package yaptranslator

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetAWSTranslatorFromEnv() (Translator, error) {

	// create translator
	keyID := os.Getenv("TEST_AWS_KEY_ID")
	keySecret := os.Getenv("TEST_AWS_KEY_SECRET")
	region := os.Getenv("TEST_AWS_KEY_REGION")

	return NewTranslator(AWS, map[string]string{
		"key":    keyID,
		"secret": keySecret,
		"region": region,
	})

}
func TestAWSCloudTranslator_GetTranslation(t *testing.T) {
	a := assert.New(t)
	c, e := GetAWSTranslatorFromEnv()
	a.NoError(e)

	tr, e := c.GetTranslation("你好", "zh", "en")
	a.NoError(e)
	a.Equal("Hello", tr)

	tr, e = c.GetTranslation("Hello", "en", "zh")
	a.NoError(e)
	a.Equal("你好", tr)

	_, e = c.GetTranslation("你不好", "not exist", "cn")
	a.Error(e) // throw error
}
