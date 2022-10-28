package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

const (
	configPath = "./config/config.yml"
)

type Config struct {
	Pg struct {
		URL string `yaml:"url"`
	} `yaml:"postgres"`

	HTTP struct {
		Port string `yaml:"port"`
	} `yaml:"http"`
}

func NewConfig() (*Config, error) {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
