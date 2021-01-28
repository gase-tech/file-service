package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/model"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/security"
	"net/http"
)

func Example(ctx *gin.Context) {
	id := security.GetCurrentCustomerId(*ctx.Request)

	customer, err := security.GetCurrentCustomer(*ctx.Request)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.BaseResponse{
			Message: err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"current-customer-id": id,
			"customer":            customer,
		})
	}
}
