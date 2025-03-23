package usecase

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"

	"github.com/danniels/shortening-url/internal/domain"
)

func (uc *usecase) ShortURL(url string) (*domain.UrlMapping, error) {
	hash := sha256.Sum256([]byte(url))

	encoded := base64.URLEncoding.EncodeToString(hash[:])

	shortURL := strings.TrimRight(encoded, "=")[:8]

	return &domain.UrlMapping{
		LongURL:  url,
		ShortURL: shortURL,
	}, nil
}
