package service

import (
	"context"
	"fmt"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/store"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"
	"io"
)

func ExistBucket(bucketName string) bool {
	ctx := context.Background()
	exists, err := store.MinioClient.BucketExists(ctx, bucketName)
	if err != nil {
		panic(err)
	}
	return exists
}

func CreateBucket(bucketName string) error {
	ctx := context.Background()
	return store.MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
}

func UploadStorage(bucketName string, fileName string, size int64, file io.Reader) (minio.UploadInfo, error) {
	ctx := context.Background()
	object, err := store.MinioClient.PutObject(ctx, bucketName, fileName, file, size, minio.PutObjectOptions{})
	if err == nil {
		log.Info().Msg(fmt.Sprint(object))
	}
	return object, err
}

func DeleteFromStorage(bucketName string, fileName string) error {
	ctx := context.Background()
	err := store.MinioClient.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})

	if err != nil {
		log.Err(err)
	}
	return err
}

func GetMedia(bucketName string, fileName string) (*minio.Object, error) {
	ctx := context.Background()

	return store.MinioClient.GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
}
