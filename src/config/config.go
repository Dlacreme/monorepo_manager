package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

const (
	FilePath = "monorepo.toml"
)

type Config struct {
	Name       string
	Workspaces []Workspace
}

type Workspace struct {
	Name   string
	Folder string
}

func Load() (*Config, error) {
	return LoadFromFile(FilePath)
}

func LoadFromFile(filepath string) (*Config, error) {
	f, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var conf Config

	_, err = toml.Decode(string(f), &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
