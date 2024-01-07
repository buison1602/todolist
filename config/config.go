package config

import (
	"fmt"
	"github.com/buison1602/todolist/web"
	"github.com/spf13/viper"
)

func LoadConfig(ConfigPath, ConfigName, ConfigType string) (*web.AppConfig, error) {
	var appConfig *web.AppConfig

	viper.AddConfigPath(ConfigPath)
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigType)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("could not read the config file: %v", err)
	}

	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal: %v", err)
	}
	return appConfig, err
}
