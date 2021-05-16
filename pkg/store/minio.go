package store

import "github.com/minio/minio-go/v7"

var MinioClient *minio.Client

type MinioConfig struct {
	URL            string `envconfig:"MINIO_URL"`
	AccessKey      string `envconfig:"MINIO_ACCESS_KEY"`
	SecretKey      string `envconfig:"MINIO_SECRET_KEY"`
	MainBucketName string `envconfig:"MINIO_BUCKET_NAME"`
	Secure         bool   `envconfig:"MINIO_SECURE"`
}
