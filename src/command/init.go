package command

import (
	"errors"
	"fmt"
	"monorepo_manager/src/global"
	"monorepo_manager/src/template"
	"os"
)

// Create a default config file
// (init keyword was already taken)
func initialize() error {
	if isFileExisting(global.ConfigFilePath) {
		return errors.New(fmt.Sprintf(global.FileExisting, global.ConfigFilePath))
	}
	os.WriteFile(global.ConfigFilePath, template.EmptyConfigFile(), 0644)

	return nil
}

func isFileExisting(filepath string) bool {
	if _, err := os.Stat(filepath); err == nil {
		return true
	}
	return false
}
