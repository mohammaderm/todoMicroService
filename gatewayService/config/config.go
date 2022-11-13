package config

import (
	"io/ioutil"
	"os"

	// _ "embed"

	"gopkg.in/yaml.v2"
)

// go:embed config.yaml
// var configs []byte

func NewConfig(path string) (*Config, error) {
	config := Config{}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
