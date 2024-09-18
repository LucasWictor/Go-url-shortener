package shortener

// URLService defines the behavior of a URL shortener service.
type URLService interface {
    ShortenURL(longURL string) string
    GetLongURL(shortURL string) (string, bool)
}
