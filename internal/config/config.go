package config

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Config is db characteristics
type Config struct {
	DB   DBConfig   `yaml:"database"`
	Http HttpConfig `yaml:"http"`
}

type DBConfig struct {
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Port   string `yaml:"port"`
	DBName string `yaml:"dbname"`
}

type HttpConfig struct {
	Port string `yaml:"port"`
}

func GetDBConfig() (DBConfig, error) {

	f, err := os.Open("internal/config/config.yaml")
	if err != nil {
		return DBConfig{}, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return DBConfig{}, err
	}

	return cfg.DB, nil
}

func GetHttpConfig() (HttpConfig, error) {
	f, err := os.Open("internal/config/config.yaml")
	if err != nil {
		return HttpConfig{}, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return HttpConfig{}, err
	}

	return cfg.Http, nil
}
