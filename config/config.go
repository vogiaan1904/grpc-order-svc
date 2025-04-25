package config

import (
	"strings"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	Mongo MongoConfig
	Grpc  GrpcMicroserviceConfig
	Log   LogConfig
}

type MongoConfig struct {
	DatabaseName string `env:"DATABASE_NAME" envDefault:"order_db"`
	DatabaseUri  string `env:"DATABASE_URI" envDefault:"mongodb://localhost:27018"`
}

type GrpcMicroserviceConfig struct {
	AuthSvcAddr    string `env:"AUTH_SERVICE_ADDRESS" envDefault:"localhost:50051"`
	ProductSvcAddr string `env:"PRODUCT_SERVICE_ADDRESS" envDefault:"localhost:50053"`
}

type LogConfig struct {
	Level        string   `env:"LOG_LEVEL" envDefault:"debug"`
	Encoding     string   `env:"LOG_ENCODING" envDefault:"development"`
	Mode         string   `env:"LOG_MODE" envDefault:"console"`
	RedactFields []string `env:"LOG_REDACT_FIELDS" envDefault:"password,token,secret"`
}

func Load() (*Config, error) {
	godotenv.Load()
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	// Process the LOG_REDACT_FIELDS env var
	if len(cfg.Log.RedactFields) == 1 && strings.Contains(cfg.Log.RedactFields[0], ",") {
		cfg.Log.RedactFields = strings.Split(cfg.Log.RedactFields[0], ",")
	}

	return cfg, nil
}
