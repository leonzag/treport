package config

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
)

func AppDataDir(appName string) (string, error) {
	if dir := os.Getenv("XDG_DATA_HOME"); dir != "" {
		return filepath.Join(dir, appName), nil
	}
	if runtime.GOOS == "linux" {
		return filepath.Join("~/.local/share", appName), nil
	}

	dir, err := os.UserConfigDir()

	return filepath.Join(dir, appName), err
}

func SQLiteDBPath(dirPath, dbName string) (string, error) {
	var perm fs.FileMode = 0755

	if err := os.MkdirAll(dirPath, perm); err != nil && !os.IsExist(err) {
		return "", fmt.Errorf("failed to prepare database directory: %w", err)
	}

	return filepath.Join(dirPath, dbName), nil
}

func SQLiteDBPathDefault(appName string) (string, error) {
	dbDir, err := AppDataDir(appName)
	if err != nil {
		return "", err
	}

	return SQLiteDBPath(dbDir, DBName)
}
