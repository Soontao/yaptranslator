package yaptranslator

import (
	"fmt"
)

// TranslatorType enum
type TranslatorType string

// Translator interface
type Translator interface {
	GetTranslation(s string, l string, t string) (string, error)
}

const (
	// ALICLOUD Translator
	ALICLOUD = TranslatorType("ALICLOUD")
)

// NewTranslator instance
func NewTranslator(provider TranslatorType, props map[string]string) (Translator, error) {

	switch provider {
	case ALICLOUD:
		return NewAliCloudTranslator(props["region"], props["key"], props["secret"])
	default:
		return nil, fmt.Errorf("Not support provider: %s", provider)
	}

}
