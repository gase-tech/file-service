package service

import (
	"errors"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/database"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/store"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"
	"io"
	"mime/multipart"
	"strconv"
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
) (uint, error) {
	if extension != "" {
		fileName += "." + extension
	}

	mediaInfo := database.MediaInfo{
		Name:        fileName,
		CustomerID:  customerID,
		TenantID:    tenantID,
		Description: description,
		Extension:   extension,
	}

	err := MediaInfoDao.Save(&mediaInfo)

	if err != nil {
		log.Err(err).Msg("Can not save media info.")
		return 0, err
	}

	_, err = UploadStorage(store.AppConfig.MinioConfig.MainBucketName, fileName, size, file)

	if err != nil {
		log.Err(err)
		MediaInfoDao.DeleteByID(mediaInfo.ID)
		return 0, err
	}

	return mediaInfo.ID, nil
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

func DeleteFile(id string) error {
	parseUint, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		log.Err(err)
		return errors.New("Id param must be numeric")
	}

	var info database.MediaInfo
	err = MediaInfoDao.FindByID(&info, uint(parseUint))

	if err != nil || info.ID == 0 {
		if err != nil {
			log.Err(err)
		}
		return errors.New("Record not found")
	}

	err = DeleteFromStorage(store.AppConfig.MinioConfig.MainBucketName, info.Name)

	if err != nil {
		return err
	}

	MediaInfoDao.DeleteByID(uint(parseUint))

	return nil
}

func PageableMediaInfos(page int, size int) interface{} {
	return MediaInfoDao.GetPageable(page, size)
}

func GetMediaInfoById(id string) (database.MediaInfo, error) {
	var info database.MediaInfo

	u64, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return info, err
	}

	err = MediaInfoDao.FindByID(&info, uint(u64))

	if err != nil {
		return info, err
	}

	return info, nil
}

func GetMediaContent(id string) (*minio.Object, database.MediaInfo, error) {
	u64, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return &minio.Object{}, database.MediaInfo{}, err
	}

	var info database.MediaInfo

	err = MediaInfoDao.FindByID(&info, uint(u64))

	if err != nil {
		return &minio.Object{}, database.MediaInfo{}, err
	}

	media, err := GetMedia(store.AppConfig.MinioConfig.MainBucketName, info.Name)

	if err != nil {
		return &minio.Object{}, database.MediaInfo{}, err
	}

	return media, info, nil
}