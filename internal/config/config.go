package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	addr string
}

type Config struct {
	Env          string `yaml:"env"  env-required:"true" env-default:"production"`
	Storage_path string `yaml:"storage_path"`
	HTTPServer   `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("no config path specified")
		}
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("error reading config file %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("error reading config file %s", err.Error())
	}

	return &cfg
}
