package config

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Config is db characteristics
type Config struct {
	DB DBConfig `yaml:"database"`
}

type DBConfig struct {
	Username string `yaml:"user"`
	Password string `yaml:"pass"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

func GetDBConfig() (DBConfig, error) {

	f, err := os.Open("config/config.yaml")
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
