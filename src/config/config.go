package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

const (
	FilePath = "monorepo.toml"
)

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

type Config struct {
	Name       string
	Workspaces []Workspace
}

func (c *Config) IsWorkspaceExisting(name string) bool {
	for _, ws := range c.Workspaces {
		if ws.Name == name {
			return true
		}
	}
	return false
}

type Workspace struct {
	Name   string
	Folder string
}
