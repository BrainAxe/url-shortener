package main

import (
	"fmt"

	"github.com/BrainAxe/url-shortener/handler"
	"github.com/BrainAxe/url-shortener/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API!🚀 ",
		})
	})

	r.POST("/api/shorten", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Note that store initialization happens here
	store.InitializeStore("redis")

	err := r.Run(":9000")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
