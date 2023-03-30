package config

import (
	"flag"
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	BindAddr string `toml:"bind_addr"`

	Technodom struct {
		DatabaseUrl string `toml:"database_url"`
		Ttl         int64  `toml:"ttl"`
	}
	LogLevel     string        `toml:"log_level"`
	TimeDuration time.Duration `toml:"time_duration"`
}

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "config/config.toml", "path to config file")
	flag.Parse()
}

func NewConfig() *TomlConfig {
	config := &TomlConfig{}
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
