package config

import (
	"log"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	AppPort          string `env:"APP_PORT" envDefault:"8080"`
	DatabaseHost     string `env:"DATABASE_HOST" envDefault:"localhost"`
	DatabasePort     string `env:"DATABASE_PORT" envDefault:"5432"`
	DatabaseName     string `env:"DATABASE_NAME" envDefault:"unc_db"`
	DatabaseUser     string `env:"DATABASE_USER" envDefault:"postgres"`
	DatabasePassword string `env:"DATABASE_PASSWORD" envDefault:"postgres"`
}

var (
	cfg  Config
	once sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		_, filename, _, _ := runtime.Caller(0)
		dir := filepath.Dir(filename)
		parentDir := filepath.Dir(dir)
		envPath := filepath.Join(parentDir, ".env")

		if err := godotenv.Load(envPath); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		if err := env.Parse(&cfg); err != nil {
			log.Fatalf("Error parsing config: %v", err)
		}
	})

	return &cfg
}
