package server

import (
	"github.com/gin-gonic/gin"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/middleware"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/store"
	"github.com/rs/zerolog/log"
)

func StartServer() {
	server := gin.New()
	/*
	server.Use(cors.Middleware(cors.Config{
		// Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE, OPTION",
		RequestHeaders: "Origin, Authorization, Content-Type, currentCustomerId, currentTenantId, instance-id, authorization",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: false,
		ValidateHeaders: false,
	}))
	 */
	server.Use(middleware.Logger(), middleware.RequestID())
	// server.Use(middleware.CORSMiddleware(), middleware.Logger(), middleware.RequestID())
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
