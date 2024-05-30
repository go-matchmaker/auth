package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

const (
	configPath = "config.yml"
)

type (
	Container struct {
		App      `yaml:"app"`
		Log      `yaml:"log"`
		HTTP     *HTTP     `yaml:"http"`
		Token    *Token    `yaml:"token"`
		PSQL     *PSQL     `yaml:"psql"`
		MONGO    *MONGO    `yaml:"mongo"`
		Settings *Settings `yaml:"settings"`
	}

	App struct {
		Name string `env-required:"true" yaml:"name" env:"APP_NAME"`
	}

	Log struct {
		Level int `env-required:"true" yaml:"level" env:"LOG_LEVEL"`
	}

	Token struct {
		TokenTTL   time.Duration `env-required:"true" yaml:"token_ttl" env:"TOKEN_TTL"`
		RefreshTTL time.Duration `env-required:"true" yaml:"refresh_ttl" env:"REFRESH_TTL"`
	}

	HTTP struct {
		Host string `env-required:"true" yaml:"host" env:"HTTP_HOST"`
		Port int    `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	PSQL struct {
		URL     string `env-required:"true" yaml:"url" env:"PSQL_URL"`
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PSQL_POOL_MAX"`
	}

	MONGO struct {
		URL            string `env-required:"true" yaml:"url" env:"MONGO_URL"`
		DBName         string `env-required:"true" yaml:"db_name" env:"MONGO_DB_NAME"`
		UserCollection string `env-required:"true" yaml:"user_collection" env:"MONGO_USER_COLLECTION"`
	}

	Settings struct {
		ServerReadTimeout int `env-required:"true" yaml:"server_read_timeout" env:"SERVER_READ_TIMEOUT"`
		DBConnAttempts    int `env-required:"true" yaml:"db_conn_attempts" env:"DB_CONN_ATTEMPTS"`
		DBConnTimeout     int `env-required:"true" yaml:"db_conn_timeout" env:"DB_CONN_TIMEOUT"`
	}
)

func NewConfig() (*Container, error) {
	cfg := new(Container)

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
