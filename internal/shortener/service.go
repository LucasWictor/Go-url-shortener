package shortener

import "math/rand"

// InMemoryURLService is a simple in-memory implementation of the URLService.
type InMemoryURLService struct {
    urlMap map[string]string
}

// NewInMemoryURLService creates a new InMemoryURLService.
func NewInMemoryURLService() *InMemoryURLService {
    return &InMemoryURLService{urlMap: make(map[string]string)}
}

// ShortenURL generates a short URL and stores it.
func (s *InMemoryURLService) ShortenURL(longURL string) string {
    shortURL := generateShortURL()
    s.urlMap[shortURL] = longURL
    return shortURL
}

// GetLongURL retrieves the original long URL from the short URL.
func (s *InMemoryURLService) GetLongURL(shortURL string) (string, bool) {
    longURL, ok := s.urlMap[shortURL]
    return longURL, ok
}

// generateShortURL generates a random 6-character short URL.
func generateShortURL() string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    shortURL := make([]byte, 6)
    for i := range shortURL {
        shortURL[i] = charset[rand.Intn(len(charset))]
    }
    return string(shortURL)
}
