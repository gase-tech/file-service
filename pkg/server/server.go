package server

import (
	"github.com/gin-gonic/gin"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/middleware"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/store"
	"github.com/rs/zerolog/log"
)

func StartServer(useCorsProtection bool) {
	server := gin.New()
	if useCorsProtection {
		server.Use(middleware.CORSMiddleware())
	}
	server.Use(middleware.Logger(), middleware.RequestID())
	server.MaxMultipartMemory = 300 << 20 // 100MiB

	setHealthHandlers(server)
	v1g := server.Group("/api/v1")
	setV1RegisterHandlers(v1g)

	err := server.Run(":" + store.AppConfig.Port)

	if err != nil {
		log.Error().Err(err)
		panic(err)
	}
}
