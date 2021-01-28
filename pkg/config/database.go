package config

import (
	"fmt"
	_const "github.com/imminoglobulin/e-commerce-backend/file-service/pkg/const"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/database"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/model"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setDBConfigForCloud(cloudConfig model.CloudConfig, cfg *store.ApplicationConfig) {
	for _, s := range cloudConfig.PropertySources {
		dbHost := s.Source["database.host"]
		if cfg.DBConfig.Host == "" && dbHost != nil && dbHost != "" {
			cfg.DBConfig.Host = dbHost.(string)
		}

		dbPort := s.Source["database.port"]
		if cfg.DBConfig.Port == 0 && dbPort != nil && dbPort != "" {
			cfg.DBConfig.Port = int(dbPort.(float64))
		}

		dbUsername := s.Source["database.username"]
		if cfg.DBConfig.Username == "" && dbUsername != nil && dbUsername != "" {
			cfg.DBConfig.Username = dbUsername.(string)
		}

		dbPassword := s.Source["database.password"]
		if cfg.DBConfig.Password == "" && dbPassword != nil && dbPassword != "" {
			cfg.DBConfig.Password = dbPassword.(string)
		}

		dbName := s.Source["database.name"]
		if cfg.DBConfig.DBName == "" && dbName != nil && dbName != "" {
			cfg.DBConfig.DBName = dbName.(string)
		}

		dbType := s.Source["database.type"]
		if cfg.DBConfig.Type == "" && dbType != nil && dbType != "" {
			cfg.DBConfig.Type = dbType.(string)
		}
	}
}

func ConnectDB(cfg store.DBConfig) {
	if cfg.Host == "" || cfg.Port == 0 || cfg.Type == "" || cfg.DBName == "" || cfg.Password == "" || cfg.Username == "" {
		err := errors.New("Host, port, type, name, username and password not null")
		log.Error().Err(err).Msg(fmt.Sprint(cfg))
		panic(err)
	}

	var connectionString string

	if cfg.Type == _const.MYSQL {
		connectionString = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.Username,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.DBName,
		)

		open, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

		if err != nil {
			log.Error().Err(err)
			panic(err)
		} else {
			store.Connection = open
		}
	}

	migrateDB()
}

func migrateDB() {
	err := store.Connection.AutoMigrate(
		&database.MediaInfo{},
	)

	if err != nil {
		log.Err(err)
		panic(err)
	}
}
