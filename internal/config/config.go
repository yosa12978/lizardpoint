package config

import (
	"os"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

var (
	c          Config
	configInit sync.Once
)

type Config struct {
	App struct {
		LogLevel string `yaml:"log_level" envconfig:"LP_LOG_LEVEL"`
	} `yaml:"app"`
	Server struct {
		Addr         string `yaml:"addr" envconfig:"LP_ADDR"`
		RootPassword string `yaml:"root_password" envconfig:"LP_ROOT_PASSWORD"`
	} `yaml:"server"`
	Postgres struct {
		User           string `yaml:"user" envconfig:"LP_POSTGRES_USER"`
		Password       string `yaml:"password" envconfig:"LP_POSTGRES_PASSWORD"`
		Database       string `yaml:"database" envconfig:"LP_POSTGRES_DATABASE"`
		SSLMode        string `yaml:"sslmode" envconfig:"LP_POSTGRES_SSLMODE"`
		URL            string `yaml:"url" envconfig:"LP_POSTGRES_URL"`
		MigrationsPath string `yaml:"migrations_path" envconfig:"LP_POSTGRES_MIGRATIONS_PATH"`
	} `yaml:"postgres"`
}

func Get() Config {
	configInit.Do(func() {
		if err := readYaml("config.yaml", &c); err != nil {
			panic(err)
		}
		if err := readEnv(&c); err != nil {
			panic(err)
		}
	})
	return c
}

func readYaml(filename string, cfg *Config) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return yaml.NewDecoder(f).Decode(cfg)
}

func readEnv(cfg *Config) error {
	return envconfig.Process("", cfg)
}
