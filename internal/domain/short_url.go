package domain

type UrlMapping struct {
	LongURL  string `json:"long_url" validate:"required"`
	ShortURL string `json:"short_url" validate:"required"`
}

type ShortUrlRequest struct {
	URL string `json:"url" validate:"required" binding:"required"`
}
