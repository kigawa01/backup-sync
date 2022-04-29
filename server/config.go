package main

import (
	"io"
	"os"
)

func defaultConf(config Config) Config {
	config.Lib = ""
	return config
}

type Config struct {
	Lib string `yaml:"lib"`
}

func load() {
	config := Config{}
	yaml, err := os.ReadFile(ConfigDir + "config.yml")
	if err == io.EOF {
		defaultConf(config)
		
	}
}
