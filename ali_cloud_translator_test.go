package yaptranslator

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAliCloudTranslator(t *testing.T) {

	// export TEST_ALICLOUD_KEY_ID=
	// export TEST_ALICLOUD_KEY_SECRET=
	// export TEST_ALICLOUD_KEY_REGION=cn-hangzhou

	a := assert.New(t)

	// create translator
	keyID := os.Getenv("TEST_ALICLOUD_KEY_ID")
	keySecret := os.Getenv("TEST_ALICLOUD_KEY_SECRET")
	region := os.Getenv("TEST_ALICLOUD_KEY_REGION")
	c, e := NewAliCloudTranslator(region, keyID, keySecret)
	a.NoError(e)

	tr, e := c.GetTranslation("你好", "zh", "en")
	a.NoError(e)
	a.Equal("Hello", tr)

	_, e = c.GetTranslation("你不好", "not exist", "cn")
	a.Error(e) // throw error
	a.Equal("translate from source to target not support", e.Error())
}
