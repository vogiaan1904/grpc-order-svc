package config

import "github.com/caarlos0/env/v9"

type Config struct {
	Mongo MongoConfig
	Grpc  GrpcMicroserviceConfig
}

type MongoConfig struct {
	DatabaseName string `env:"DATABASE_NAME" envDefault:"order_db"`
	DatabaseUri  string `env:"DATABASE_URI" envDefault:"mongodb://localhost:27017"`
}

type GrpcMicroserviceConfig struct {
	AuthServiceAddress    string `env:"AUTH_SERVICE_ADDRESS" envDefault:"localhost:50051"`
	ProductServiceAddress string `env:"PRODUCT_SERVICE_ADDRESS" envDefault:"localhost:50053"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
