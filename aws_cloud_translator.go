package yaptranslator

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
)

// AWSCloudTranslator provider
type AWSCloudTranslator struct {
	c *translate.Translate
}

// NewAWSCloudTranslator instance
func NewAWSCloudTranslator(region, keyID, keySecret string) (*AWSCloudTranslator, error) {

	sess, err := session.NewSession(
		&aws.Config{
			Region:      &region,
			Credentials: credentials.NewStaticCredentials(keyID, keySecret, ""),
		},
	)

	if err != nil {
		return nil, err
	}

	t := translate.New(sess)

	return &AWSCloudTranslator{c: t}, nil
}

// GetTranslation text
func (t *AWSCloudTranslator) GetTranslation(s string, l string, tl string) (r string, e error) {

	res, err := t.c.Text(&translate.TextInput{
		SourceLanguageCode: &l,
		TargetLanguageCode: &tl,
		Text:               &s,
	})

	if err != nil {
		return "", err
	}

	return *res.TranslatedText, nil
}
