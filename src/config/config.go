package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

const (
	FilePath   = "monorepo.toml"
	envVarName = "MM_WORKSPACE"
)

func Load() (*Config, error) {
	return LoadFromFile(FilePath)
}

func LoadFromFile(filepath string) (*Config, error) {
	f, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var conf Config = Config{
		EnvVarName: envVarName,
	}

	_, err = toml.Decode(string(f), &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

type Config struct {
	EnvVarName string
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

func (c *Config) GetWorkspace(name string) *Workspace {
	for _, ws := range c.Workspaces {
		if ws.Name == name {
			return &ws
		}
	}
	return nil
}

func (c *Config) GetRoot() *Workspace {
	for _, ws := range c.Workspaces {
		if ws.IsRoot {
			return &ws
		}
	}
	return nil
}

type Workspace struct {
	Name   string
	Folder string
	IsRoot bool
}
