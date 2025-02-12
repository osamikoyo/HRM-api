package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

const CONFIG_PATH = "config.yml"

type Config struct{
	VocationSrcUrl string `yaml:"vocation_service_addr"`
	WorkerSvcUrl string `yaml:"worker_service_addr"`
	Port uint64 `yaml:"port"`
	Host string `yaml:"host"`
}

func LoadConfig() *Config {
	defaultConfig := &Config{
		VocationSrcUrl: "localhost:50051",
			WorkerSvcUrl: "localhost:50052",
			Port: 8080,
			Host: "localhost",
	}

	file, err := os.Open(CONFIG_PATH)
	if err != nil{
		fmt.Print(err)

		return defaultConfig
	}

	data, err := io.ReadAll(file)
	if err != nil{
		fmt.Println(err)

		return defaultConfig
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg);err != nil{
		fmt.Println(err)

		return defaultConfig
	}

	return &cfg
}