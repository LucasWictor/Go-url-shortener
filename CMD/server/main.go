package main

import (
	"go-url-shortener/internal/shortener"  
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create the URL shortener service
	urlService := shortener.NewInMemoryURLService()

	r := gin.Default()

	// routes
	r.POST("/shorten", func(c *gin.Context) {
		shortener.HandleShortenURL(c, urlService)
	})
	r.GET("/:shortURL", func(c *gin.Context) {
		shortener.HandleRedirectURL(c, urlService)
	})

	r.Run(":8080")
}
