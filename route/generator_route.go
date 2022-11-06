package route

import (
	"github.com/amdtaufiq/sorturl-GO/app/handler"
	"github.com/amdtaufiq/sorturl-GO/app/service"
	"github.com/gin-gonic/gin"
)

func GeneratorRoute(routerGroup *gin.RouterGroup) {
	generatorService := service.GeneratorService()
	generatorHandler := handler.GeneratorHandler(generatorService)

	routerGroup.POST("/short-url", generatorHandler.CreateGerateSortURL)
	routerGroup.GET("/short-url/:url", generatorHandler.HandleShortURLRedirect)
}
