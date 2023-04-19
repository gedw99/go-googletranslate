package translate

import (
	"go-googletranslate/languages"
)

type Translator interface {
	Translate(text string, sourceLanguage languages.LanguageCode, targetLanguage languages.LanguageCode) (string, error)
	SetToCache(text, translatedText string, sourceLanguage languages.LanguageCode, targetLanguage languages.LanguageCode) error
	GetFromCache(text string, sourceLanguage languages.LanguageCode, targetLanguage languages.LanguageCode) (string, error)
}
