package config

import (
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/model"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/service"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/store"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

func setMinioConfigForCloud(cloudConfig model.CloudConfig, cfg *store.ApplicationConfig) {
	for _, s := range cloudConfig.PropertySources {
		url := s.Source["minio.url"]
		if cfg.MinioConfig.URL == "" && url != nil && url != "" {
			cfg.MinioConfig.URL = url.(string)
		}

		accessKey := s.Source["minio.access-key"]
		if cfg.MinioConfig.AccessKey == "" && accessKey != nil && accessKey != "" {
			cfg.MinioConfig.AccessKey = accessKey.(string)
		}

		secretKey := s.Source["minio.secret-key"]
		if cfg.MinioConfig.SecretKey == "" && secretKey != nil && secretKey != "" {
			cfg.MinioConfig.SecretKey = secretKey.(string)
		}

		bucketName := s.Source["minio.bucket-name"]
		if cfg.MinioConfig.MainBucketName == "" && bucketName != nil && bucketName != "" {
			cfg.MinioConfig.MainBucketName = bucketName.(string)
		}
	}
}

func ConnectMinio(cfg store.MinioConfig) {
	client, err := minio.New(cfg.URL, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Err(err).Msg("Can not create minio client.")
		panic(err)
	} else {
		log.Info().Msg("Connected minio..." + cfg.URL)
	}

	store.MinioClient = client

	if !service.ExistBucket(cfg.MainBucketName) {
		err := service.CreateBucket(cfg.MainBucketName)
		if err != nil {
			panic(err)
		}
	}
}
