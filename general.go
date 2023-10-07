package pexels

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
	// supported locales are:
	// en-US, pt-BR, es-ES, ca-ES, de-DE, it-IT, fr-FR, sv-SE, id-ID, pl-PL,
	// ja-JP, zh-TW, zh-CN, ko-KR, th-TH, nl-NL, hu-HU, vi-VN, cs-CZ, da-DK,
	// fi-FI, uk-UA, el-GR, ro-RO, nb-NO, sk-SK, tr-TR, ru-RU.
	Locale string `query:"locale"`
	// supported orientations are: landscape, portrait, square.
	Orientation string `query:"orientation"`
	// The minimum size, supported sizes are: small, medium, large.
	Size string `query:"size"`
}
