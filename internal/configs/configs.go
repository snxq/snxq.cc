package configs

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config is config for all apps
type Config struct {
	Mysql MySQL `yaml:"mysql"`
}

// MySQL is config for mysql database
type MySQL struct {
	DSN          string        `yaml:"dsn"`
	MaxIdleConns int           `yaml:"max_idle_conns"`
	MaxOpenConns int           `yaml:"max_open_conns"`
	MaxLifeTime  time.Duration `yaml:"max_life_time"`
}

// Load loads and parses the config at path
func Load(path string) (*Config, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if stat.IsDir() {
		return nil, fmt.Errorf("config cannot be a dir - %s", path)
	}

	var cfg Config
	if err := parseYamlFile(path, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseYamlFile(path string, out interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(b, out); err != nil {
		return fmt.Errorf("error unmarshaling %s: %+v", path, err)
	}
	return nil
}
