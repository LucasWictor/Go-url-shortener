package shortener

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// HandleShortenURL handles the URL shortening request
func HandleShortenURL(c *gin.Context, service URLService) {
    longURL := c.PostForm("url")
    if longURL == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
        return
    }

    shortURL := service.ShortenURL(longURL)
    c.JSON(http.StatusOK, gin.H{"shortURL": shortURL})
}

// HandleRedirectURL handles the redirection for the short URL
func HandleRedirectURL(c *gin.Context, service URLService) {
    shortURL := c.Param("shortURL")
    longURL, ok := service.GetLongURL(shortURL)
    if !ok {
        c.JSON(http.StatusNotFound, gin.H{"message": "Short URL not found"})
        return
    }

    c.Redirect(http.StatusMovedPermanently, longURL)
}
