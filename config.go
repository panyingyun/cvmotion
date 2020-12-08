package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	DeviceID string `yaml:"deviceid"`
	Host     string `yaml:"host"`
	Prefix   string `yaml:"prefix"`
}

var config Config

func InitConfig() error {
	yamlFile, err := ioutil.ReadFile("prod.yaml")
	if err != nil {
		panic("config not found")
	}
	err = yaml.Unmarshal(yamlFile, &config)
	return err
}
