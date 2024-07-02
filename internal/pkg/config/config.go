package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type HTTP struct {
	Host string `yaml:"127.0.0.1"`
	Port int    `yaml:"port"`
}

type Log struct {
	Level string `yaml:"level"`
}

type DB struct {
	Dirver string `yaml:"driver"`
	DSN    string `yaml:"dsn"`
}

type Config struct {
	Env  string `yaml:"env"`
	HTTP HTTP   `yaml:"http"`
	Log  Log    `yaml:"log"`
	DB   DB     `yaml:"db"`
}

func ParseWithPath(path string) (*Config, error) {
	buf, err := readConfig(path)
	if err != nil {
		return nil, err
	}

	return ParseWithBytes(buf)
}

func ParseWithBytes(buf []byte) (*Config, error) {
	return Parse(buf)
}

func Parse(buf []byte) (*Config, error) {
	cfg := Config{}
	if err := yaml.Unmarshal(buf, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func readConfig(path string) ([]byte, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("配置文件 %s 是空的", path)
	}

	return data, nil
}
