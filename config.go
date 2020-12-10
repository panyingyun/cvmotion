package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type CameraConfig struct {
	DeviceID string `yaml:"deviceid"`
}

type WebConfig struct {
	Enable bool   `yaml:"enable"`
	Host   string `yaml:"host"`
	Port   string `yaml:"post"`
}

type VideoConfig struct {
	Enable bool    `yaml:"enable"`
	Fps    float64 `yaml:"fps"`
	Prefix string  `yaml:"prefix"`
}

type WindowConfig struct {
	Enable bool   `yaml:"enable"`
	Title  string `yaml:"title"`
	Width  int    `yaml:"width"`
	Height int    `yaml:"height"`
}

type Config struct {
	Camera CameraConfig `yaml:"camera"`
	Web    WebConfig    `yaml:"web"`
	Video  VideoConfig  `yaml:"video"`
	Window WindowConfig `yaml:"window"`
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
