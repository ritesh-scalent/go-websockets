package config

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

type configs struct {
	Redis   RedisConfig   `yaml:"redis"`
	Service ServiceConfig `yaml:"service"`
}

var Configs configs

func Init(Config, ConfigPath *string) {
	var configPath string
	if Config == nil || *Config == "dev" {
		_, b, _, _ := runtime.Caller(0)
		BasePath := filepath.Dir(b)
		configPath = BasePath + "/configs.yaml"
	} else {
		configPath = *ConfigPath
	}
	load(configPath)
}

func load(ConfigsPath string) {
	yamlFile, err := ioutil.ReadFile(ConfigsPath)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &Configs)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
