package config

import (
	"monorepo_manager/src/test"
	"os"
	"testing"
)

var testFilePath string = "test_monorepo.toml"

func TestLoad(t *testing.T) {
	t.Run("load the config file", func(t *testing.T) {
		createConfig(defaultConfig())
		subject, err := LoadFromFile(testFilePath)
		test.AssertNoError(t, err)
		test.AssertStringEq(t, subject.Name, "default")
		test.AssertNumEq(t, len(subject.Workspaces), 1)
		test.AssertStringEq(t, subject.EnvVarName, "MM_WORKSPACE")
	})

	t.Run("return an error if the file is missing", func(t *testing.T) {
		cleanTestFile()
		_, err := LoadFromFile(testFilePath)
		test.AssertError(t, err)
	})

	t.Run("return an error if the file is invalid", func(t *testing.T) {
		createConfig(invalidConfig())
		_, err := LoadFromFile(testFilePath)
		test.AssertError(t, err)
		cleanTestFile()
	})
}

func createConfig(config []byte) {
	cleanTestFile()
	os.WriteFile(testFilePath, config, 0644)
}

func defaultConfig() []byte {
	return []byte(`name = "default"

[[workspaces]]
name = "workspace1"
folder = "workspace1_folder"
`)
}

func invalidConfig() []byte {
	return []byte(`invalid config file!`)
}

func cleanTestFile() {
	os.Remove(testFilePath)
}
