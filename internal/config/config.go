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
	Server struct {
		Addr         string `yaml:"addr" envconfig:"LP_ADDR"`
		RootPassword string `yaml:"root_pass" envconfig:"LP_ROOT_PASS"`
	} `yaml:"server"`
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
