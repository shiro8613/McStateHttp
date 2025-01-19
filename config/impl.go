package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

var internalConfig Config

func Load(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(b, &internalConfig); err != nil {
		return err
	}

	return nil
}

func CreateConfig(path string) error {
	conf := &Config{}
	conf.NewDefault()

	b, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}

	fp, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = fp.Write(b)
	
	return err
}

func GetConfig() Config {
	return internalConfig
}