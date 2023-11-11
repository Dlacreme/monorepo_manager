package command

import (
	"fmt"
	"monorepo_manager/src/config"
	"monorepo_manager/src/content"
	"monorepo_manager/src/test"
	"os"
	"testing"
)

func TestInitialize(t *testing.T) {
	t.Run("create the config file as expected", func(t *testing.T) {
		os.Remove(config.FilePath)
		err := initialize()
		test.AssertNoError(t, err)
		os.Remove(config.FilePath)
	})

	t.Run("return an error if the file is already existing", func(t *testing.T) {
		os.WriteFile(config.FilePath, []byte{'0'}, 0644)
		err := initialize()
		test.AssertErrorEq(t, err, fmt.Sprint(config.FilePath, " ", content.FileExisting))
		os.Remove(config.FilePath)
	})
}
