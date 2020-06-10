package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config wraps the configuration structs for the various application parts
type Config struct {
	Database struct {
		Name string `default:"instagram_bot.db"`
	}
	Server struct {
		Host string `default:"localhost"`
		Port string `default:"3000"`
	}
}

// Parse ...
func Parse(envFile string) *Config {

	var err error
	if fileExists(envFile) {
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Error parsing configuration file: %s", err)
		}

	}

	var cfg Config

	// with config file loaded into env values, we can now parse env into our config struct
	err = envconfig.Process("app", &cfg)
	if err != nil {
		log.Fatalf("Error parsing configuration from environment: %s", err)
	}

	return &cfg
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
