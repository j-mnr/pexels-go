package pexels

type Pagination struct {
	TotalResults uint32 `json:"total_results"`
	Page         uint16 `json:"page"`
	PerPage      uint8  `json:"per_page"`  // Default: 15, Max: 80
	PrevPage     string `json:"prev_page"` // Optional
	NextPage     string `json:"next_page"` // Optional
}
