package pexels

// Pagination is a common response struct for many endpoints that details how
// to go to the previous and next pages as well as the total number of results,
// what page you are on currently and how many results per page.
type Pagination struct {
	TotalResults uint32 `json:"total_results"`
	Page         uint16 `json:"page"`
	PerPage      uint8  `json:"per_page"`  // Default: 15, Max: 80
	PrevPage     string `json:"prev_page"` // Optional
	NextPage     string `json:"next_page"` // Optional
}

// Orientation is the way the photo or video is oriented.
type Orientation string

const (
	// Landscape is oriented horizontally
	Landscape Orientation = "landscape"
	// Portrait is oriented vertically
	Portrait Orientation = "portrait"
	// Square is oriented symetrically
	Square Orientation = "square"
)

// Size is the minimum video or photo size
type Size string

const (
	// Large 24MP for photos or 4K for videos
	Large Size = "large"
	// Medium 12MP for photos or Full HD for videos
	Medium Size = "medium"
	// Small 4MP for photos or HD for videos
	Small Size = "small"
)

// Locale is the locale of the search you are performing.
type Locale string

// All of the Locales
const (
	EnUS Locale = "en-US"
	PtBR Locale = "pt-BR"
	EsES Locale = "es-ES"
	CaES Locale = "ca-ES"
	DeDE Locale = "de-DE"
	ItIT Locale = "it-IT"
	FrFR Locale = "fr-FR"
	SvSE Locale = "sv-SE"
	IdID Locale = "id-ID"
	PlPL Locale = "pl-PL"
	JaJP Locale = "ja-JP"
	ZhTW Locale = "zh-TW"
	ZhCN Locale = "zh-CN"
	KoKR Locale = "ko-KR"
	ThTH Locale = "th-TH"
	NlNL Locale = "nl-NL"
	HuHU Locale = "hu-HU"
	ViVN Locale = "vi-VN"
	CsCZ Locale = "cs-CZ"
	DaDK Locale = "da-DK"
	FiFI Locale = "fi-FI"
	UkUA Locale = "uk-UA"
	ElGR Locale = "el-GR"
	RoRO Locale = "ro-RO"
	NbNO Locale = "nb-NO"
	SkSK Locale = "sk-SK"
	TrTR Locale = "tr-TR"
	RuRU Locale = "ru-RU"
)
