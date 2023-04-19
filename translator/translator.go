package translator

import (
	"go-googletranslate/googletranslate"
	"go-googletranslate/translate"
)

const (
	defaultCacheFileName = "go-googletranslate.cache.json"
	defaultCacheMaxSize  = 100
)

type Translator = translate.Translator
type GoogleTranslate = googletranslate.GoogleTranslate

func NewGoogleTranslate(cacheFileName string, cacheMaxSize int) (*GoogleTranslate, error) {

	if cacheFileName == "" {
		cacheFileName = defaultCacheFileName
	}

	if cacheMaxSize == 0 {
		cacheMaxSize = defaultCacheMaxSize
	}

	return googletranslate.NewGoogleTranslate(cacheFileName, cacheMaxSize)
}

func NewDefaultGoogleTranslate() (*GoogleTranslate, error) {
	return NewGoogleTranslate("", 0)
}
