package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	Mongo MongoConfig
	Grpc  GrpcMicroserviceConfig
}

type MongoConfig struct {
	DatabaseName string `env:"DATABASE_NAME" envDefault:"order_db"`
	DatabaseUri  string `env:"DATABASE_URI" envDefault:"mongodb://localhost:27018"`
}

type GrpcMicroserviceConfig struct {
	AuthServiceAddress    string `env:"AUTH_SERVICE_ADDRESS" envDefault:"localhost:50051"`
	ProductServiceAddress string `env:"PRODUCT_SERVICE_ADDRESS" envDefault:"localhost:50053"`
}

func Load() (*Config, error) {
	godotenv.Load()
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
