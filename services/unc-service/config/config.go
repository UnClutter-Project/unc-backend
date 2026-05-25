package config

import (
	"log"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	AppPort                   string        `env:"APP_PORT" envDefault:"8080"`
	CorsAllowedOrigin         string        `env:"CORS_ALLOWED_ORIGIN" envDefault:""`
	JWTSecret                 string        `env:"JWT_SECRET" envDefault:""`
	DatabaseHost              string        `env:"DATABASE_HOST" envDefault:"localhost"`
	DatabasePort              string        `env:"DATABASE_PORT" envDefault:"5432"`
	DatabaseName              string        `env:"DATABASE_NAME" envDefault:"unc_db"`
	DatabaseUser              string        `env:"DATABASE_USER" envDefault:"postgres"`
	DatabasePassword          string        `env:"DATABASE_PASSWORD" envDefault:"postgres"`
	BrevoAPIKey               string        `env:"BREVO_API_KEY" envDefault:""`
	ClientEndpoint            string        `env:"CLIENT_ENDPOINT" envDefault:"http://localhost:3000"`
	ClientEmailVerifyPath     string        `env:"CLIENT_EMAIL_VERIFY_PATH" envDefault:"/verify"`
	CloudflareID              string        `env:"CLOUDFLARE_ID" envDefault:""`
	ObjectStorageAccessID     string        `env:"OBJECT_STORAGE_ACCESS_ID" envDefault:""`
	ObjectStorageAccessSecret string        `env:"OBJECT_STORAGE_ACCESS_SECRET" envDefault:""`
	BucketName                string        `env:"BUCKET_NAME" envDefault:"unc"`
	PresignDuration           time.Duration `env:"PRESIGN_DURATION" envDefault:"15m"`
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
			log.Printf("Warning error loading .env file: %v", err)
		}

		if err := env.Parse(&cfg); err != nil {
			log.Fatalf("Error parsing config: %v", err)
		}
	})

	return &cfg
}
