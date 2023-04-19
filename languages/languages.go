package languages

type LanguageCode int

const (
	AutoDetect LanguageCode = iota
	AF
	SQ
	AM
	AR
	HY
	AZ
	EU
	BE
	BN
	BS
	BG
	CA
	CEB
	NY
	ZH_CN
	ZH_TW
	CO
	HR
	CS
	DA
	NL
	EN
	EO
	ET
	FIL
	FI
	FR
	FY
	GL
	KA
	DE
	EL
	GU
	HT
	HA
	HAW
	IW
	HI
	HMONG
	HU
	IS
	IG
	ID
	GA
	IT
	JA
	JV
	KN
	KK
	KM
	KO
	KU
	KY
	LO
	LA
	LV
	LT
	LB
	MK
	MG
	MS
	ML
	MT
	MI
	MR
	MN
	MY
	NE
	NO
	PS
	FA
	PL
	PT
	PA
	RO
	RU
	SM
	GD
	SR
	ST
	SN
	SD
	SI
	SK
	SL
	SO
	ES
	SU
	SW
	SV
	TG
	TA
	TT
	TE
	TH
	TR
	TK
	UK
	UR
	UZ
	VI
	CY
	XHOSA
	YI
	YO
	ZU
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
