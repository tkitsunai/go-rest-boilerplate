package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var (
	_config Config
	once    = sync.Once{}
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

func setupConfig() Config {
	v := viper.New()

	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath("internal/config")
	v.AddConfigPath("/")
	v.AddConfigPath(".")

	v.AllowEmptyEnv(true)
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	err = v.Unmarshal(&_config)
	if err != nil {
		fmt.Println(err)
	}

	return _config
}

func Get() *Config {
	once.Do(func() {
		_config = setupConfig()
	})
	return &_config
}
