package main

import (
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/config"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/server"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/store"
)

func main() {
	config.PrepareAppConfig()
	server.StartServer(!store.AppConfig.IsConnectServiceRegistry)
}
