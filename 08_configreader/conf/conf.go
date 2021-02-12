package conf

type ConfigType struct {
	port uint
	db_url, jaeger_url, sentry_url, kafka_broker, some_app_id, some_app_key string
}
//ReadFlags reads config from command-line flags
func ReadFlags(config *ConfigType) {}

//ReadEnv reads data from Environment variables
func ReadEnv(config *ConfigType) {}
