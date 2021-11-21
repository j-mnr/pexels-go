package pexels

type Locale string

const (
	EN_US Locale = "en-US"
	PT_BR Locale = "pt-BR"
	ES_ES Locale = "es-ES"
	CA_ES Locale = "ca-ES"
	DE_DE Locale = "de-DE"
	IT_IT Locale = "it-IT"
	FR_FR Locale = "fr-FR"
	SV_SE Locale = "sv-SE"
	ID_ID Locale = "id-ID"
	PL_PL Locale = "pl-PL"
	JA_JP Locale = "ja-JP"
	ZH_TW Locale = "zh-TW"
	ZH_CN Locale = "zh-CN"
	KO_KR Locale = "ko-KR"
	TH_TH Locale = "th-TH"
	NL_NL Locale = "nl-NL"
	HU_HU Locale = "hu-HU"
	VI_VN Locale = "vi-VN"
	CS_CZ Locale = "cs-CZ"
	DA_DK Locale = "da-DK"
	FI_FI Locale = "fi-FI"
	UK_UA Locale = "uk-UA"
	EL_GR Locale = "el-GR"
	RO_RO Locale = "ro-RO"
	NB_NO Locale = "nb-NO"
	SK_SK Locale = "sk-SK"
	TR_TR Locale = "tr-TR"
	RU_RU Locale = "ru-RU"
)

type Orientation string

const (
	Landscape Orientation = "landscape"
	Portrait  Orientation = "portrait"
	Square    Orientation = "square"
)

type Size string

const (
	Large  Size = "large"
	Medium Size = "medium"
	Small  Size = "small"
)
