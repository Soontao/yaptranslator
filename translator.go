package yaptranslator

import "golang.org/x/text/language"

// Translator interface
type Translator interface {
	GetTranslation(s string, l language.Tag, t language.Tag)
}
