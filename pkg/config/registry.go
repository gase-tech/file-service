package config

import (
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/discovery"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/model"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/store"
)

func setRegistryConfigForCloud(cloudConfig model.CloudConfig, cfg *store.ApplicationConfig)  {
	for _, s := range cloudConfig.PropertySources {
		url := s.Source["registry.url"]
		if cfg.RegistryConfig.URL == "" && url != nil && url != "" {
			cfg.RegistryConfig.URL = url.(string)
		}

		username := s.Source["registry.username"]
		if cfg.RegistryConfig.Username == "" && username != nil && username != "" {
			cfg.RegistryConfig.Username = username.(string)
		}

		password := s.Source["registry.password"]
		if cfg.RegistryConfig.Password == "" && password != nil && password != "" {
			cfg.RegistryConfig.Password = password.(string)
		}

		useSSL := s.Source["registry.use-ssl"]
		if useSSL != nil && useSSL != "" {
			cfg.RegistryConfig.UseSSL = useSSL.(bool)
		}
	}
}

func ServiceRegister(cfg store.RegistryConfig)  {
	config := discovery.RegistrationVariables{ServiceRegistryURL: cfg.URL, UserName: cfg.Username, Password: cfg.Password}
	go discovery.ManageDiscovery(config)
}
