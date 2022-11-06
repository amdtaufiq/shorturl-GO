package service

import (
	"testing"

	"github.com/amdtaufiq/sorturl-GO/app/dto"
	"github.com/stretchr/testify/assert"
)

func TestShortLinkGenerator(t *testing.T) {
	var generatorService generatorService
	longURL := "https://github.com/amdtaufiq/TodoList-GO"
	shortURL := generatorService.GenerateSortUrl(dto.CreateGeratorRequest{
		LongURL: longURL,
	})

	assert.Equal(t, shortURL.SortURL, "KuETbg")
}
