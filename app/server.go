package app

import (
	"log"

	"github.com/amdtaufiq/sorturl-GO/app/config"
	"github.com/amdtaufiq/sorturl-GO/app/store"
	"github.com/amdtaufiq/sorturl-GO/route"
)

func Run() {
	// Router
	router := config.InitRouter("debug")
	v1 := router.Group("/v1")

	// Routes
	route.GeneratorRoute(v1)

	store.InitializeStore()

	err := router.Run()
	if err != nil {
		log.Fatalf("can't run server")
	}
}
