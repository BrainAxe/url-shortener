package main

import (
	"fmt"
	"net/http"
	"os"

	handler "github.com/BrainAxe/url-shortener/handlers"
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
	r.LoadHTMLGlob("templates/*")
	// Serve static files
	r.Static("/static", "./static")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.POST("/api/shorten", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/api/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Note that store initialization happens here
	store.InitializeStore("redis")

	err := r.Run(":" + os.Getenv("HOST_PORT"))
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
