package googletranslate

type LanguageCode int

const (
	AutoDetect LanguageCode = iota
	AFRIKAANS
	ALBANIAN
	AMHARIC
	ARABIC
	ARMENIAN
	AZERBAIJANI
	BASQUE
	BELARUSIAN
	BENGALI
	BOSNIAN
	BULGARIAN
	CATALAN
	CEBUANO
	CHICHEWA
	CHINESE
	CHINESE_TW
	CORSICAN
	CROATIAN
	CZECH
	DANISH
	DUTCH
	ENGLISH
	ESPERANTO
	ESTONIAN
	FILIPINO
	FINNISH
	FRENCH
	FRISIAN
	GALICIAN
	GEORGIAN
	GERMAN
	GREEK
	GUJARATI
	HAITIAN_CREOLE
	HAUSA
	HAWAIIAN
	HEBREW
	HINDI
	HMONG
	HUNGARIAN
	ICELANDIC
	IGBO
	INDONESIAN
	IRISH
	ITALIAN
	JAPANESE
	JAVANESE
	KANNADA
	KAZAKH
	KHMER
	KOREAN
	KURDISH
	KYRGYZ
	LAO
	LATIN
	LATVIAN
	LITHUANIAN
	LUXEMBOURGISH
	MACEDONIAN
	MALAGASY
	MALAY
	MALAYALAM
	MALTESE
	MAORI
	MARATHI
	MONGOLIAN
	BURMESE
	NEPALI
	NORWEGIAN
	PASHTO
	PERSIAN
	POLISH
	PORTUGUESE
	PUNJABI
	ROMANIAN
	RUSSIAN
	SAMOAN
	SCOTS_GAELIC
	SERBIAN
	SESOTHO
	SHONA
	SINDHI
	SINHALA
	SLOVAK
	SLOVENIAN
	SOMALI
	SPANISH
	SUNDANESE
	SWAHILI
	SWEDISH
	TAJIK
	TAMIL
	TATAR
	TELUGU
	THAI
	TURKISH
	TURKMEN
	UKRAINIAN
	URDU
	UZBEK
	VIETNAMESE
	WELSH
	XHOSA
	YIDDISH
	YORUBA
	ZULU
)

func (lc LanguageCode) String() string {
	return [...]string{
		"auto", "af", "sq", "am", "ar", "hy", "az", "eu", "be", "bn", "bs", "bg", "ca", "ceb", "ny",
		"zh-CN", "zh-TW", "co", "hr", "cs", "da", "nl", "en", "eo", "et", "fil", "fi", "fr",
		"fy", "gl", "ka", "de", "el", "gu", "ht", "ha", "haw", "iw", "hi", "hmn", "hu",
		"is", "ig", "id", "ga", "it", "ja", "jv", "kn", "kk", "km", "ko", "ku", "ky",
		"lo", "la", "lv", "lt", "lb", "mk", "mg", "ms", "ml", "mt", "mi", "mr", "mn",
		"my", "ne", "no", "ps", "fa", "pl", "pt", "pa", "ro", "ru", "sm", "gd", "sr",
		"st", "sn", "sd", "si", "sk", "sl", "so", "es", "su", "sw", "sv", "tg", "ta",
		"tt", "te", "th", "tr", "tk", "uk", "ur", "uz", "vi", "cy", "xh", "yi", "yo", "zu",
	}[lc]
}
