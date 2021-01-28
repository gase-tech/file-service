package server

import (
	"github.com/gin-gonic/gin"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/api"
	v1 "github.com/imminoglobulin/e-commerce-backend/file-service/pkg/api/v1"
)

func setHealthHandlers(e *gin.Engine) {
	e.GET("/", api.Home)
	e.GET("/status", api.Status)
	e.GET("/healthcheck", api.HealthCheck)
}

func setV1RegisterHandlers(g *gin.RouterGroup) {
	g.POST("/media/upload", v1.UploadFile)
	g.DELETE("/media/:id", v1.DeleteFile)
	g.GET("/media/info-id/:id", v1.GetMediaInfoById)
	g.GET("/media/info/pageable", v1.GetMediaInfoPageable)
	g.GET("/media/content/:id", v1.GetMediaContent)
	g.GET("/", v1.Example)
}
