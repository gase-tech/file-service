package store

import "strings"

type RegistryConfig struct {
	URL      string `envconfig:"REGISTRY_URL"`
	Username string `envconfig:"REGISTRY_USERNAME"`
	Password string `envconfig:"REGISTRY_PASSWORD"`
	UseSSL   bool   `envconfig:"USE_SSL"`
}

type ServiceInfos struct {
	Services []ServiceApp
}

type ServiceApp struct {
	Name             string
	InstanceBaseUrls []string
}

func (si ServiceInfos) GetServiceUrl(serviceName string) string {
	for _, s := range si.Services {
		if strings.EqualFold(s.Name, serviceName) {
			if len(s.InstanceBaseUrls) > 0 {
				return s.InstanceBaseUrls[0]
			}
		}
	}
	return ""
}

var RegisteredServices *ServiceInfos
