package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Token string `yaml:"token"`
}

//Reads config file and parses it into Config struct
func ReadConfig(path string) Config {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var c Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
	return c
}
