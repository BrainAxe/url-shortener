package handler

import (
	"net/http"

	"github.com/BrainAxe/url-shortener/store"
	"github.com/BrainAxe/url-shortener/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	err := c.ShouldBindJSON(&creationRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	strRandom := uuid.New()
	shortUrl := utils.GenerateShortLink(creationRequest.LongUrl, strRandom.String())
	store.StoreService.Strategy.SaveUrlMapping(shortUrl, creationRequest.LongUrl)

	host := c.Request.Host + "/api/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})

}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")

	initialUrl := store.StoreService.Strategy.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
