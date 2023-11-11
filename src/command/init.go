package command

import (
	"errors"
	"fmt"
	"monorepo_manager/src/config"
	"monorepo_manager/src/content"
	"os"
)

// Create a default config file
// (init keyword was already taken)
func initialize() error {
	if isFileExisting(config.FilePath) {
		return errors.New(fmt.Sprint(config.FilePath, " ", content.FileExisting))
	}

	err := os.WriteFile(config.FilePath, config.EmptyConfigFile(), 0644)
	if err != nil {
		return err
	}

	fmt.Println(config.FilePath, content.FileCreated)
	return nil
}

func isFileExisting(filepath string) bool {
	if _, err := os.Stat(filepath); err == nil {
		return true
	}
	return false
}
