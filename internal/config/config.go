package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"addr" env:"HTTP_ADDR" env-default:"localhost:8080"`
}

type Config struct {
	Env          string `yaml:"env"  env-required:"true" env-default:"production"`
	Storage_path string `yaml:"storage_path"`
	HTTPServer   `yaml:"http_server"`
}

// CONFIG_PATH environment variable
//         ↓
// If empty → check -config command-line flag
//         ↓
// If still empty → program stops
//         ↓
// Check whether file exists
//         ↓
// cleanenv reads YAML
//         ↓
// Store values inside Config
//         ↓
// Return Config

func MustLoad() *Config {

	// Create a variable to store the path of the config file.
	var configPath string

	// First, try to get the config file path from the CONFIG_PATH environment variable.
	configPath = os.Getenv("CONFIG_PATH")

	// If CONFIG_PATH is empty, try to get the path from the command-line flag.
	if configPath == "" {

		// Create a command-line flag called "config".
		// Example: go run main.go -config=config.yaml
		flags := flag.String("config", "", "path to config file")

		// Read the command-line arguments.
		flag.Parse()

		// Get the value provided to the -config flag.
		configPath = *flags

		// If the config path is still empty, stop the program.
		if configPath == "" {

			// Print the error and terminate the program.
			log.Fatal("no config path specified")
		}
	}

	// Check whether the config file exists.
	if _, err := os.Stat(configPath); err != nil {

		// If the file does not exist or cannot be accessed, stop the program.
		log.Fatalf("error reading config file %s", configPath)
	}

	// Create an empty Config struct.
	var cfg Config

	// Read the config file and put its values into the cfg struct.
	err := cleanenv.ReadConfig(configPath, &cfg)

	// Check if there was an error while reading the config file.
	if err != nil {

		// Print the error and terminate the program.
		log.Fatalf("error reading config file %s", err.Error())
	}

	// Return the address of the cfg struct.
	return &cfg

	// cfg.Env = "production"
	// cfg.StoragePath = "./storage"
	// cfg.Addr = "localhost:8080"
}
