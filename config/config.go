package config

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config struct
type Config struct {
	Port        int    `env:"PORT" envDefault:"4000"`
	Auth        bool   `env:"AUTH"`
	TokenSecret string `env:"TOKEN_SECRET"`
	DB          string `env:"DB"`
	DBUser      string `env:"DB_USER"`
	DBPass      string `env:"DB_PASS"`
	DBMock      bool   `env:"DB_MOCK"`
}

// Init environment variables
func Init() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	configFile := fmt.Sprint("config/.env.", appEnv)
	log.Println("Using config file:", configFile)

	err := godotenv.Load(configFile)
	if err != nil {
		log.Fatal(err)
	}
}

// Get configs
func Get() Config {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
