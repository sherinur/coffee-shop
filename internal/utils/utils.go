package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// FileExists checks if a file or directory exists at the specified path.
// It returns true if the file or directory exists, false if it does not exist,
// and an error if there is an issue with checking the file status.
func FileExists(path string) (bool, error) {
	// Use os.Stat to check the file status
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// CreateFile creates a file at the given path. It ensures the directory exists
// and creates the file if it does not already exist. Returns an error if something
// goes wrong during directory creation, file creation, or any other operation.
func CreateFile(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory '%s': %w", dir, err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file '%s': %w", path, err)
	}
	defer file.Close()

	return nil
}

// CreateDir creates a directory at the specified path if it does not exist.
// Returns an error if there is an issue checking the path or creating the directory.
func CreateDir(path string) error {
	exists, err := FileExists(path)
	if err != nil {
		return fmt.Errorf("failed to check if directory exists: %w", err)
	}

	if exists {
		return nil
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory '%s': %w", path, err)
	}

	return nil
}

// RemoveDir removes the directory at the given path and its contents.
// Returns an error if the path is empty or if an error occurs during removal.
func RemoveDir(path string) error {
	if path == "" {
		return fmt.Errorf("filepath is empty: cannot remove directory")
	}

	if err := os.RemoveAll(path); err != nil {
		return fmt.Errorf("failed to remove directory at '%s': %w", path, err)
	}

	return nil
}

// RemoveFile deletes the file at the given path.
// Returns an error if the file path is empty or if an error occurs during file removal.
func RemoveFile(path string) error {
	if path == "" {
		return fmt.Errorf("filepath is empty: cannot remove file")
	}

	if err := os.Remove(path); err != nil {
		return fmt.Errorf("failed to remove file at '%s': %w", path, err)
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

// DirExists() checks if the given directory path exists.
// It returns nil if the directory exists, or an error if it does not exist or another error occurs.
func DirExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("directory does not exist: %s", path)
	}
	if err != nil {
		return err
	}

	return nil
}

// FileEmpty() checks if the given file is empty.
// It returns true if the file size is 0, and false otherwise.
func FileEmpty(file *os.File) bool {
	if stat, err := file.Stat(); err != nil {
		return false
	} else if stat.Size() == 0 {
		return true
	}

	return false
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
  --cfg S      Path to the config file.`)
}

// ValidatePort checks if the provided port string is a valid number
// and within the valid range (1024 to 65535).
// It returns an error if the port is invalid.
func ValidatePort(port string) error {
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("invalid port: '%s' is not a valid number", port)
	}

	if portNum < 1024 || portNum > 65535 {
		return fmt.Errorf("invalid port: '%d' must be between 1024 and 65535", portNum)
	}

	return nil
}

// ValidateDirName() checks the name of a single directory element.
func ValidateDirName(name string) error {
	validName := regexp.MustCompile(`^[a-z0-9_-]+$`)

	if !validName.MatchString(name) {
		return errors.New("invalid directory name: use only lowercase letters, digits, hyphens, and underscores")
	}

	for _, r := range name {
		if r > 127 {
			return errors.New("invalid characters in directory name: use ASCII (Latin) characters only")
		}
	}

	return nil
}

// ValidatePath checks the entire path by validating each directory element.
func ValidatePath(path string) error {
	cleanPath := filepath.Clean(path)

	dirs := strings.Split(cleanPath, "/")
	for _, dir := range dirs {
		if dir == "" {
			continue
		}

		if err := ValidateDirName(dir); err != nil {
			return fmt.Errorf("invalid path element '%s': %w", dir, err)
		}
	}

	return nil
}

// ValidateDir checks if the directory exists and creates it if necessary.
func ValidateDir(dir string) error {
	if err := ValidatePath(dir); err != nil {
		return fmt.Errorf("path validation error: %w", err)
	}

	info, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
			return nil
		}
		return fmt.Errorf("directory check error: %w", err)
	}

	if !info.IsDir() {
		return errors.New("the specified path is not a directory")
	}

	return nil
}
