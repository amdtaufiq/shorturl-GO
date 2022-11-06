package handler

import (
	"net/http"

	"github.com/amdtaufiq/sorturl-GO/app/dto"
	"github.com/amdtaufiq/sorturl-GO/app/service"
	"github.com/amdtaufiq/sorturl-GO/app/store"
	"github.com/gin-gonic/gin"
)

type generatorHandler struct {
	generatorService service.IGeneratorService
}

func GeneratorHandler(generatorService service.IGeneratorService) *generatorHandler {
	return &generatorHandler{generatorService}
}

func (h *generatorHandler) CreateGerateSortURL(ctx *gin.Context) {
	var createGeratorRequest dto.CreateGeratorRequest
	if err := ctx.ShouldBindJSON(&createGeratorRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := h.generatorService.GenerateSortUrl(createGeratorRequest)
	store.SaveUrlMapping(response.SortURL, createGeratorRequest.LongURL)
	response.Message = "Short URL createad"
	response.SortURL = "http://localhost:8080/v1/short-url/" + response.SortURL
	ctx.JSON(http.StatusOK, response)
}

func (h *generatorHandler) HandleShortURLRedirect(ctx *gin.Context) {
	var inputSortURL dto.GetSortURLRequest

	if err := ctx.ShouldBindUri(&inputSortURL); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initialUrl := store.RetrieveInitialUrl(inputSortURL.SortURL)
	ctx.Redirect(302, initialUrl)
}
