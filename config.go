package main

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Notifies []Notify `yaml:"notifies"`
}

type Notify struct {
	Files []GithubActionFile `yaml:"files"`
}

type GithubActionFile struct {
	Name      string `yaml:"name"`
	Type      string `yaml:"type"`
	Added     bool   `yaml:"added"`
	Modified  bool   `yaml:"modified"`
	Removed   bool   `yaml:"removed"`
	Message   string `yaml:"message"`
	Messenger `yaml:"messenger"`
}

type Messenger struct {
	Type string `yaml:"type"`
	Url  string `yaml:"url"`
}

func NewConfig(filename string) (*Config, error) {
	var config Config

	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("file is not readable")
	}

	err = yaml.Unmarshal([]byte(yamlFile), &config)
	if err != nil {
		return nil, errors.New("yaml file is not invalid")
	}

	return &config, nil
}
