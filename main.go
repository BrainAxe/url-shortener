package main

import (
	"fmt"
	"os"

	"github.com/BrainAxe/url-shortener/handlers"
	"github.com/BrainAxe/url-shortener/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Load env file
	errENV := godotenv.Load()
	if errENV != nil {
		panic(fmt.Sprintf("Error loading .env file - Error: %v", errENV))
	}

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

	err := r.Run(":" + os.Getenv("HOST_PORT"))
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
