package store

import "gorm.io/gorm"

var Connection *gorm.DB

type DBConfig struct {
	Host						string				`envconfig:"DB_HOST"`
	Port						int					`envconfig:"DB_PORT"`
	Username					string				`envconfig:"DB_USERNAME"`
	Password					string				`envconfig:"DB_PASSWORD"`
	DBName						string				`envconfig:"DB_NAME"`
	Type						string				`envconfig:"DB_TYPE"`
}