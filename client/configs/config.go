package configs

import "github.com/spf13/viper"

type conf struct {


	Title  							string `mapstructure:"TITLE"`
	BackgroundColor 				string `mapstructure:"BACKGROUND_COLOR"`
	ResponseTime    				string `mapstructure:"RESPONSE_TIME"`
	ExternalCallURL 				string `mapstructure:"EXTERNAL_CALL_URL"`
	ExternalCallMethod 				string `mapstructure:"EXTERNAL_CALL_METHOD"`
	RequestNameOTEL 				string `mapstructure:"REQUEST_NAME_OTEL"`
	OTEL_SERVICE_NAME 				string `mapstructure:"OTEL_SERVICE_NAME"`
	OTEL_EXPORTER_OTLP_ENDPOINT 	string `mapstructure:"OTEL_EXPORTER_OTLP_ENDPOINT"`
	HTTP_PORT  						string `mapstructure:"HTTP_PORT"`
}



func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile("config.env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
