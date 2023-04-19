package googletranslate

const (
	defaultCacheFileName = "go-googletranslate.cache.json"
	defaultCacheMaxSize  = 100
)

type Translator interface {
	Translate(text string, sourceLanguage LanguageCode, targetLanguage LanguageCode) (string, error)
	SetToCache(text, translatedText string, sourceLanguage LanguageCode, targetLanguage LanguageCode) error
	GetFromCache(text string, sourceLanguage LanguageCode, targetLanguage LanguageCode) (string, error)
}

func NewGoogleTranslator(cacheFileName string, cacheMaxSize int) (*GoogleTranslate, error) {

	if cacheFileName == "" {
		cacheFileName = defaultCacheFileName
	}

	if cacheMaxSize == 0 {
		cacheMaxSize = defaultCacheMaxSize
	}

	return NewGoogleTranslate(cacheFileName, cacheMaxSize)
}

func NewDefaultGoogleTranslator() (*GoogleTranslate, error) {
	return NewGoogleTranslator("", 0)
}
