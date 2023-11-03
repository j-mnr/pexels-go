package pexels

// Locale is an enum; all of them start with "Locale".
type Locale interface {
	locale()
}

type locale string

func (locale) locale() {}

const (
	LocaleEN_US locale = "en-US"
	LocalePT_BR locale = "pt-BR"
	LocaleES_ES locale = "es-ES"
	LocaleCA_ES locale = "ca-ES"
	LocaleDE_DE locale = "de-DE"
	LocaleIT_IT locale = "it-IT"
	LocaleFR_FR locale = "fr-FR"
	LocaleSV_SE locale = "sv-SE"
	LocaleID_ID locale = "id-ID"
	LocalePL_PL locale = "pl-PL"
	LocaleJA_JP locale = "ja-JP"
	LocaleZH_TW locale = "zh-TW"
	LocaleZH_CN locale = "zh-CN"
	LocaleKO_KR locale = "ko-KR"
	LocaleTH_TH locale = "th-TH"
	LocaleNL_NL locale = "nl-NL"
	LocaleHU_HU locale = "hu-HU"
	LocaleVI_VN locale = "vi-VN"
	LocaleCS_CZ locale = "cs-CZ"
	LocaleDA_DK locale = "da-DK"
	LocaleFI_FI locale = "fi-FI"
	LocaleUK_UA locale = "uk-UA"
	LocaleEL_GR locale = "el-GR"
	LocaleRO_RO locale = "ro-RO"
	LocaleNB_NO locale = "nb-NO"
	LocaleSK_SK locale = "sk-SK"
	LocaleTR_TR locale = "tr-TR"
	LocaleRU_RU locale = "ru-RU"
)

// Size is an enum; all of them start with "Size".
type Size interface {
	size()
}

type size string

func (size) size() {}

const (
	SizeSmall  size = "small"
	SizeMedium size = "medium"
	SizeLarge  size = "large"
)

// Orientation is an enum; all of them start with "Orientation".
type Orientation interface {
	orientation()
}

type orientation string

func (orientation) orientation() {}

const (
	OrientationLandscape = "landscape"
	OrientationPortrait  = "portrait"
	OrientationSquare    = "square"
)

// Pagination is a common response struct for many endpoints that details how
// to go to the previous and next pages as well as the total number of results,
// what page you are on currently and how many results per page.
//
//nolint:tagliatelle
type Pagination struct {
	TotalResults uint32 `json:"total_results"`
	Page         uint16 `json:"page"`
	PerPage      uint8  `json:"per_page"` // Default: 15, Max: 80

	PrevPage string `json:"prev_page"`
	NextPage string `json:"next_page"`
}

// General is the common parameters found between video and photo queries.
type General struct {
	Locale      Locale      `query:"locale"`
	Orientation Orientation `query:"orientation"`
	Size        Size        `query:"size"`
}
