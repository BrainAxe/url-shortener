package tests

import (
	"fmt"
	"testing"

	"github.com/BrainAxe/url-shortener/store"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var testStoreService = &store.StorageService{}

func init() {
	//Load env file
	errENV := godotenv.Load("../.env")
	if errENV != nil {
		panic(fmt.Sprintf("Error loading .env file - Error: %v", errENV))
	}
	testStoreService = store.InitializeStore("mongo")
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://github.com/faif/python-patterns"
	shortURL := "Jsz4k57oAX"

	// Persist data mapping
	testStoreService.Strategy.SaveUrlMapping(shortURL, initialLink)

	// Retrieve initial URL
	retrieveUrl := testStoreService.Strategy.RetrieveInitialUrl(shortURL)

	assert.Equal(t, initialLink, retrieveUrl)
}
