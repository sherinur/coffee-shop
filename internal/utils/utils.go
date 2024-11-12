package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

// FileExists checks if a file or directory exists at the specified path.
// It returns true if the file or directory exists, false if it does not exist,
// and an error if there is an issue with checking the file status.
func FileExists(path string) (bool, error) {
	// Use os.Stat to check the file status
	_, err := os.Stat(path)
	if err == nil {
		// File or directory exists
		return true, nil
	}

	// If the error is 'not exist', return false
	if os.IsNotExist(err) {
		return false, nil
	}

	// For other errors, return the error
	return false, err
}

// CreateFile creates a file at the given path. It ensures the directory exists
// and creates the file if it does not already exist. Returns an error if something
// goes wrong during directory creation, file creation, or any other operation.
func CreateFile(path string) error {
	// Ensure the directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory '%s': %w", dir, err)
	}

	// Create the file
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
	// Check if the directory already exists
	exists, err := FileExists(path)
	if err != nil {
		return fmt.Errorf("failed to check if directory exists: %w", err)
	}

	// If directory exists, no need to create it, just return nil
	if exists {
		return nil
	}

	// Attempt to create the directory and any necessary parent directories
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory '%s': %w", path, err)
	}

	return nil
}

// RemoveDir removes the directory at the given path and its contents.
// Returns an error if the path is empty or if an error occurs during removal.
func RemoveDir(path string) error {
	// Check if the path is empty and return an appropriate error
	if path == "" {
		return fmt.Errorf("filepath is empty: cannot remove directory")
	}

	// Attempt to remove the directory and its contents, and wrap any error that occurs
	if err := os.RemoveAll(path); err != nil {
		return fmt.Errorf("failed to remove directory at '%s': %w", path, err)
	}

	// Return nil if the directory was successfully removed
	return nil
}

// RemoveFile deletes the file at the given path.
// Returns an error if the file path is empty or if an error occurs during file removal.
func RemoveFile(path string) error {
	// Check if the path is empty and return an appropriate error
	if path == "" {
		return fmt.Errorf("filepath is empty: cannot remove file")
	}

	// Attempt to remove the file and wrap any error that occurs
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("failed to remove file at '%s': %w", path, err)
	}

	// Return nil if file was successfully removed
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
	// Use os.Stat to check if the directory exists
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// If the directory doesn't exist, return an error indicating that
		return fmt.Errorf("directory does not exist: %s", path)
	}
	if err != nil {
		// If another error occurred (e.g., permission issues), return that error
		return err
	}

	// Return nil if the directory exists and no errors occurred
	return nil
}

// FileEmpty() checks if the given file is empty.
// It returns true if the file size is 0, and false otherwise.
func FileEmpty(file *os.File) bool {
	// Get the file's statistics, including its size
	if stat, err := file.Stat(); err != nil {
		// Handle potential error when obtaining file stats
		// In a real-world application, you would likely want to log the error or return it
		return false // In case of error, assume the file isn't empty
	} else if stat.Size() == 0 {
		// If the file size is 0, return true indicating the file is empty
		return true
	}

	// Return false if the file is not empty
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
	// Convert the port string to an integer
	portNum, err := strconv.Atoi(port)
	if err != nil {
		// If the conversion fails, return a specific error message
		return fmt.Errorf("invalid port: '%s' is not a valid number", port)
	}

	// Check if the port number is within the valid range (1024 - 65535)
	if portNum < 1024 || portNum > 65535 {
		// Return a specific error message if the port is out of range
		return fmt.Errorf("invalid port: '%d' must be between 1024 and 65535", portNum)
	}

	// Return nil if the port is valid
	return nil
}

// ValidateDirName() checks the name of a single directory element.
func ValidateDirName(name string) error {
	// Regular expression to validate directory names:
	// Only lowercase letters, digits, hyphens, and underscores are allowed.
	validName := regexp.MustCompile(`^[a-z0-9_-]+$`)

	if !validName.MatchString(name) {
		return errors.New("invalid directory name: use only lowercase letters, digits, hyphens, and underscores")
	}

	// Check for ASCII characters only (avoid UTF-8 symbols like Cyrillic).
	for _, r := range name {
		if r > 127 {
			return errors.New("invalid characters in directory name: use ASCII (Latin) characters only")
		}
	}

	return nil
}

// ValidatePath checks the entire path by validating each directory element.
func ValidatePath(path string) error {
	// Normalize the path according to the current operating system.
	cleanPath := filepath.Clean(path)

	// Split the path into its components (directories).
	dirs := filepath.SplitList(cleanPath)
	for _, dir := range dirs {
		if dir == "" {
			continue // Skip empty elements (e.g., in the root path).
		}
		// Validate each directory name.
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

	// Check if the directory already exists.
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

	// Ensure the path is a directory and not a file.
	if !info.IsDir() {
		return errors.New("the specified path is not a directory")
	}

	return nil
}
