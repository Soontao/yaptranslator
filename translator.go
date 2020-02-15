package yaptranslator

// Translator interface
type Translator interface {
	GetTranslation(s string, l string, t string) (string, error)
}
