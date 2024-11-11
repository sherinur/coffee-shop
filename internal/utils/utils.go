package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func CreateFile(path string) error {
	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func CreateDir(path string) error {
	exists, err := FileExists(path)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}

	return nil
}

func RemoveDir(path string) error {
	if path == "" {
		return fmt.Errorf("filepath to the directory is empty")
	}

	if err := os.RemoveAll(path); err != nil {
		return fmt.Errorf("error of RemoveDir: %w", err)
	}

	return nil
}

func RemoveFile(path string) error {
	if path == "" {
		return fmt.Errorf("filepath to the file is empty")
	}

	if err := os.Remove(path); err != nil {
		return fmt.Errorf("error of RemoveFile: %w", err)
	}

	return nil
}

// DirEmpty() takes path of the directory as an argument.
// If directory empty, returns true and nil. False and nil in other case.
// Returns false and error, if error occurs.
func IsDirEmpty(path string) (bool, error) {
	isDirExits, err := FileExists(path)
	if err != nil {
		return false, err
	}

	if !isDirExits {
		return false, fmt.Errorf(path+" %w", "ErrDirNotExist")
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return false, err
	}

	return len(entries) == 0, nil
}

// GetExecPath() returns path of executable file
// Returns error, if error occurs.
func GetExecPath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("error of GetExecPath: %w", err)
	}
	return execPath, nil
}

// RemoveValue() takes slice of string and index as argument
// and removes the element in this index
func RemoveValue(records [][]string, index int) [][]string {
	if index < 0 || index >= len(records) {
		return records
	}

	return append(records[:index], records[index+1:]...)
}

func DirExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

func CustomUsage() {
	fmt.Println(`Coffee Shop Management System

Usage:
  hot-coffee [--port <N>] [--dir <S>] [--cfg <S>]
  hot-coffee --help

Options:
  --help       Show this screen.
  --port N     Port number.
  --dir S      Path to the data directory.
- --cfg S    Path to the config file`)
}

func ValidatePort(port string) error {
	// Convert the port string to an integer
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return errors.New("invalid port: not a number")
	}

	// Check if the port is in the valid range
	if portNum < 1024 || portNum > 65535 {
		return errors.New("invalid port: must be between 1024 and 65535")
	}

	return nil
}

func ValidateDir(dir string) error {
	err := CreateDir(dir)
	// Check if the path exists
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return errors.New("directory does not exist")
	}
	if err != nil {
		return err
	}

	// Check if the path is a directory
	if !info.IsDir() {
		return errors.New("path is not a directory")
	}

	return nil
}
