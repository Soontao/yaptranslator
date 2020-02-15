package yaptranslator

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
