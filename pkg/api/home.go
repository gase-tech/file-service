package api

import (
	"github.com/gin-gonic/gin"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/model"
	"net/http"
)

func Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.BaseResponse{
		Message: "UP",
	})
}

func Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.BaseResponse{
		Message: "UP",
	})
}

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.BaseResponse{
		Message: "UP",
	})
}