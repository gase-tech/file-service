package store

var AppConfig *ApplicationConfig

type ApplicationConfig struct {
	ApplicationName				string				`envconfig:"APPLICATION_NAME" default:"file-service"`
	Profile						string 				`envconfig:"PROFILE" default:"dev"`
	Port						string				`envconfig:"PORT" default:"5050"`
	IsConnectConfigServer		bool				`envconfig:"IS_CONNECT_CONFIG_SERVER" default:"true"`
	ConfigServerUrl				string				`envconfig:"CONFIG_SERVER_URL" default:"http://localhost:8888/file-service/"`
	DBConfig					*DBConfig
	IsConnectServiceRegistry	bool				`envconfig:"IS_CONNECT_SERVICE_REGISTRY" default:"true"`
	RegistryConfig				*RegistryConfig
	MinioConfig					*MinioConfig
}
