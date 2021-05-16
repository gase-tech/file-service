package service

import (
	"errors"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/database"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/store"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"io"
	"mime/multipart"
	"strings"
)

var MediaInfoDao = new(database.MediaInfoDao)

func UploadFile(
	description string,
	fileName string,
	file io.Reader,
	extension string,
	size int64,
	customerID string,
	tenantID string,
) (string, error) {
	if extension != "" {
		fileName += "." + extension
	}

	mediaInfo := database.MediaInfo{
		ID:          uuid.NewV4(),
		Name:        fileName,
		CustomerID:  customerID,
		TenantID:    tenantID,
		Description: description,
		Extension:   extension,
	}

	err := MediaInfoDao.Save(&mediaInfo)

	if err != nil {
		log.Err(err).Msg("Can not save media info.")
		return "", err
	}

	_, err = UploadStorage(store.AppConfig.MinioConfig.MainBucketName, fileName, size, file)

	if err != nil {
		log.Err(err)
		MediaInfoDao.DeleteByID(mediaInfo.ID)
		return "", err
	}

	return mediaInfo.ID.String(), nil
}

func FindFileExtension(file multipart.FileHeader) string {
	filename := file.Filename

	splitFilename := strings.Split(filename, ".")

	if len(splitFilename) > 1 {
		return splitFilename[len(splitFilename)-1]
	} else {
		return ""
	}
}

func DeleteFile(id uuid.UUID) error {
	var info database.MediaInfo
	err := MediaInfoDao.FindByID(&info, id)

	if err != nil {
		log.Err(err)
		return errors.New("Record not found")
	}

	err = DeleteFromStorage(store.AppConfig.MinioConfig.MainBucketName, info.Name)

	if err != nil {
		return err
	}

	MediaInfoDao.DeleteByID(id)

	return nil
}

func PageableMediaInfos(page int, size int) interface{} {
	return MediaInfoDao.GetPageable(page, size)
}

func GetMediaInfoById(id uuid.UUID) (database.MediaInfo, error) {
	var info database.MediaInfo

	err := MediaInfoDao.FindByID(&info, id)

	if err != nil {
		return info, err
	}

	return info, nil
}

func GetMediaContent(id uuid.UUID) (*minio.Object, database.MediaInfo, error) {
	var info database.MediaInfo

	err := MediaInfoDao.FindByID(&info, id)

	if err != nil {
		return &minio.Object{}, database.MediaInfo{}, err
	}

	media, err := GetMedia(store.AppConfig.MinioConfig.MainBucketName, info.Name)

	if err != nil {
		return &minio.Object{}, database.MediaInfo{}, err
	}

	return media, info, nil
}
