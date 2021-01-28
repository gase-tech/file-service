package main

import (
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/config"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/server"
)

func main() {
	config.PrepareAppConfig()
	server.StartServer()
}