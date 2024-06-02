package tests

import (
	"testing"

	"github.com/BrainAxe/url-shortener/store"
	"github.com/stretchr/testify/assert"
)

var testStoreService = &store.StorageService{}

func init() {
	testStoreService = store.InitializeStore("redis")
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
