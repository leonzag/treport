package report

import (
	"errors"
	"fmt"
	"os"
)

func checkDestFolder(dest string) error {
	stat, err := os.Stat(dest)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("folder doesn't exist")
	}
	if !stat.IsDir() {
		return fmt.Errorf("expected a folder but found something else")
	}
	return nil
}
