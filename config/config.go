package config

import (
	"strings"
	"time"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	Mongo    MongoConfig
	Grpc     GrpcMicroserviceConfig
	Log      LogConfig
	Temporal TemporalConfig
	Server   Server
	Http     Http
}

type Server struct {
	Port              string        `env:"PORT" envDefault:"50054"`
	Development       bool          `env:"DEVELOPMENT" envDefault:"true"`
	Timeout           time.Duration `env:"TIMEOUT" envDefault:"10s"`
	ReadTimeout       time.Duration `env:"READ_TIMEOUT" envDefault:"10s"`
	WriteTimeout      time.Duration `env:"WRITE_TIMEOUT" envDefault:"10s"`
	MaxConnectionIdle time.Duration `env:"MAX_CONNECTION_IDLE" envDefault:"10s"`
	MaxConnectionAge  time.Duration `env:"MAX_CONNECTION_AGE" envDefault:"10s"`
}

type Http struct {
	Port              string        `env:"PORT" envDefault:"50054"`
	PprofPort         string        `env:"PPROF_PORT" envDefault:"50055"`
	Timeout           time.Duration `env:"TIMEOUT" envDefault:"10s"`
	ReadTimeout       time.Duration `env:"READ_TIMEOUT" envDefault:"10s"`
	WriteTimeout      time.Duration `env:"WRITE_TIMEOUT" envDefault:"10s"`
	CookieLifeTime    int           `env:"COOKIE_LIFETIME" envDefault:"10"`
	SessionCookieName string        `env:"SESSION_COOKIE_NAME" envDefault:"session"`
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

type TemporalConfig struct {
	HostPort  string `env:"TEMPORAL_HOST_PORT" envDefault:"localhost:7233"`
	Namespace string `env:"TEMPORAL_NAMESPACE" envDefault:"default"`
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
