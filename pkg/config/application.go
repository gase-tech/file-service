package config

import (
	"encoding/json"
	eureka_client "github.com/bilalkocoglu/eureka-client"
	_extconfig "github.com/bilalkocoglu/eureka-client/config"
	_extmodel "github.com/bilalkocoglu/eureka-client/model"
	_extstore "github.com/bilalkocoglu/eureka-client/store"
	_const "github.com/imminoglobulin/e-commerce-backend/file-service/pkg/const"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/model"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/store"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strconv"
)

func PrepareAppConfig() {
	cfg := new(store.ApplicationConfig)
	registryCfg := new(_extstore.RegistryConfig)
	err := envconfig.Process(_const.CONFIG_PREFIX, cfg)

	if err != nil {
		log.Err(err).Msg("Prepare config error.")
		panic(err)
	}

	err = envconfig.Process(_const.CONFIG_PREFIX, registryCfg)

	if err != nil {
		log.Err(err).Msg("Prepare config error.")
		panic(err)
	}

	if cfg.IsConnectConfigServer && cfg.ConfigServerUrl == "" {
		panic("ConfigServerURL must be not empty.")
	}

	if cfg.IsConnectConfigServer {
		ReceiveConfigFile(cfg, registryCfg)
	}
	store.AppConfig = cfg

	ConnectDB(*cfg.DBConfig)
	ConnectMinio(*cfg.MinioConfig)

	if cfg.IsConnectServiceRegistry {
		registryCfg.AppPort = cfg.Port
		registryCfg.AppName = cfg.ApplicationName
		eureka_client.ServiceRegister(*registryCfg)
	}
}

func ReceiveConfigFile(cfg *store.ApplicationConfig, registryCfg *_extstore.RegistryConfig) {
	configServerUrl := cfg.ConfigServerUrl + cfg.Profile

	log.Print("Config server url: " + configServerUrl)
	response, err := http.Get(configServerUrl)

	if err != nil {
		log.Err(err).Msg("Config server client error.")
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Err(err).Msg("Receive response body error.")
		panic(err)
	}

	var cloudConfig model.CloudConfig
	err = json.Unmarshal(body, &cloudConfig)

	if err != nil {
		log.Err(err).Msg("Convert config server response error.")
		panic(err)
	}

	for _, s := range cloudConfig.PropertySources {
		port := s.Source["port"]
		if port != nil && port != "" {
			cfg.Port = strconv.FormatFloat(port.(float64), 'g', -1, 32)
		}
	}

	setDBConfigForCloud(cloudConfig, cfg)

	var propertySources []_extmodel.CloudConfigPropertySource
	for _, source := range cloudConfig.PropertySources {
		propertySources = append(propertySources, _extmodel.CloudConfigPropertySource{
			Name:   source.Name,
			Source: source.Source,
		})
	}
	_extconfig.SetRegistryConfigForCloud(_extmodel.CloudConfig{
		Name:            cloudConfig.Name,
		Label:           cloudConfig.Label,
		Profiles:        cloudConfig.Profiles,
		State:           cloudConfig.State,
		Version:         cloudConfig.Version,
		PropertySources: propertySources,
	}, registryCfg)
	setMinioConfigForCloud(cloudConfig, cfg)
}
