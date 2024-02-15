package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	var bodydata interface{}
	json.Unmarshal(bytes, &bodydata)

	result := make(map[string]interface{})

	switch v := bodydata.(type) {
	case map[string]interface{}:
		for key, val := range v {
			result[key] = val
		}
	}

	return result
}

func (c *Config) constructRequest() (*http.Request, error) {
	body := c.readBody()

	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.URL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

func printRequest(req *http.Request) {
	fmt.Println("Request Method:", req.Method)
	fmt.Println("Request URL:", req.URL.String())
	fmt.Println("Request Headers:")
	for key, values := range req.Header {
		for _, value := range values {
			fmt.Printf("\t%s: %s\n", key, value)
		}
	}
	fmt.Println("Request Body:")
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	fmt.Println(buf.String())
}
