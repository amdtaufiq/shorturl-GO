package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	_store := InitializeStore()
	testStoreService = _store
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	longURL := "https://github.com/amdtaufiq/TodoList-GO"
	shortURL := "KuETbg"

	SaveUrlMapping(shortURL, longURL)

	assert.Equal(t, longURL, RetrieveInitialUrl(shortURL))
}
