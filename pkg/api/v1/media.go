package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/model"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/security"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/service"
	"net/http"
	"strconv"
)

func UploadFile(ctx *gin.Context) {
	filename := ctx.PostForm("filename")
	description := ctx.PostForm("description")
	file, err := ctx.FormFile("file")

	if err != nil || filename == "" || description == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("get form err: %s", err.Error()),
		})
		return
	}

	extension := service.FindFileExtension(*file)
	openedFile, err := file.Open()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "File error",
		})
	}
	customerID := security.GetCurrentCustomerId(*ctx.Request)
	tenantID := security.GetCurrentTenantId(*ctx.Request)
	uploadInfo, err := service.UploadFile(description, filename, openedFile, extension, file.Size, customerID, tenantID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("get form err: %s", err.Error()),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"info": uploadInfo,
		})
	}
}

func DeleteFile(ctx *gin.Context) {
	id := ctx.Param("id")

	err := service.DeleteFile(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.BaseResponse{
			Message: err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, model.BaseResponse{
			Message: "Success",
		})
		return
	}
}

func GetMediaInfoPageable(ctx *gin.Context) {
	sizeStr, sizeExist := ctx.GetQuery("size")
	sizeInt := 10
	if sizeExist {
		parseInt, err := strconv.ParseInt(sizeStr, 10, 32)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.BaseResponse{Message: "Size must be numeric."})
			return
		}
		sizeInt = int(parseInt)
	}

	pageStr, pageExist := ctx.GetQuery("page")
	pageInt := 0
	if pageExist {
		parseInt, err := strconv.ParseInt(pageStr, 10, 32)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.BaseResponse{Message: "Page must be numeric."})
			return
		}
		pageInt = int(parseInt)
	}

	page := service.PageableMediaInfos(pageInt, sizeInt)

	ctx.JSON(http.StatusOK, page)
	return
}

func GetMediaInfoById(ctx *gin.Context) {
	id := ctx.Param("id")

	byId, err := service.GetMediaInfoById(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.BaseResponse{
			Message: err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, byId)
		return
	}
}

func GetMediaContent(ctx *gin.Context) {
	id := ctx.Param("id")

	object, info, err := service.GetMediaContent(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.BaseResponse{
			Message: err.Error(),
		})
		return
	} else {
		extraHeaders := map[string]string{
			"Content-Disposition": fmt.Sprintf(`attachment; filename="%s"`, info.Name),
		}

		stat, err := object.Stat()

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.BaseResponse{
				Message: err.Error(),
			})
			return
		}

		ctx.DataFromReader(http.StatusOK, stat.Size, stat.ContentType, object, extraHeaders)
		return
	}
}