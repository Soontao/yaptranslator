package yaptranslator

import (
	"errors"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alimt"
	"github.com/tidwall/gjson"
)

// AliCloudTranslator provider
type AliCloudTranslator struct {
	client *alimt.Client
}

// GetTranslation text
func (t *AliCloudTranslator) GetTranslation(s string, l string, tl string) (r string, e error) {
	req := alimt.CreateTranslateGeneralRequest()
	req.Method = "POST"
	req.FormatType = "text"
	req.SourceLanguage = l
	req.TargetLanguage = tl
	req.SourceText = s
	res, err := t.client.TranslateGeneral(req)

	if err != nil {
		return "", err
	}

	if res.Code == 200 {
		return gjson.Get(res.GetHttpContentString(), "Data.Translated").String(), nil
	}

	// not success
	return "", errors.New(res.Message)
}

// NewAliCloudTranslator instance
func NewAliCloudTranslator(region, keyID, keySecret string) (*AliCloudTranslator, error) {
	c, err := alimt.NewClientWithAccessKey(region, keyID, keySecret)
	if err != nil {
		return nil, err
	}
	return &AliCloudTranslator{client: c}, nil
}
