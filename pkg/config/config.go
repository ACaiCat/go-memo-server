package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {
	PostgreConfig `yaml:"postgre_config"`
	RedisConfig   `yaml:"redis_config"`
	ServerConfig  `yaml:"server_config"`
}

func GetConfig() *Config {
	once.Do(initConfig)
	return cfg
}

func initConfig() {
	const cfgPath = "./configs/conf.yaml"
	content, err := os.ReadFile(cfgPath)
	if err != nil {
		panic(err)
	}
	cfg = new(Config)
	if err = yaml.Unmarshal(content, cfg); err != nil {
		panic(err)
	}

}
