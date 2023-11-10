package command

import (
	"fmt"
	"monorepo_manager/src/global"
	"monorepo_manager/src/test"
	"os"
	"testing"
)

func TestInitialize(t *testing.T) {
	t.Run("create the config file as expected", func(t *testing.T) {
		os.Remove(global.ConfigFilePath)
		err := initialize()
		test.AssertNoError(t, err)
		os.Remove(global.ConfigFilePath)
	})

	t.Run("return an error if the file is already existing", func(t *testing.T) {
		os.WriteFile(global.ConfigFilePath, []byte{'0'}, 0644)
		err := initialize()
		test.AssertErrorEq(t, err, fmt.Sprintf(global.FileExisting, global.ConfigFilePath))
		os.Remove(global.ConfigFilePath)
	})
}
