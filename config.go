package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Name     string            `toml:"name"`
	Method   string            `toml:"method"`
	URL      string            `toml:"URL"`
	Headers  map[string]string `toml:"headers"`
	Bodypath string            `toml:"bodypath"`
}

func readConfigFile(filename string) Config {
	var config Config

	filepath := "Configs/" + filename

	_, err := toml.DecodeFile(filepath, &config)

	if err != nil {
		panic("nooo")
	}

	return config
}

func (c *Config) readBody() map[string]interface{} {
	filepath := "Payloads/" + c.Bodypath
	jsonFile, err := os.Open(filepath)

	if err != nil {
		panic("nooo")
	}

	defer jsonFile.Close()

	bytes, _ := io.ReadAll(jsonFile)

	var bodydata map[string]interface{}
	json.Unmarshal(bytes, &bodydata)

	return bodydata
}
