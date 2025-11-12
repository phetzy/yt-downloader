package utils

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// ExpandHomeDir expands ~ to the user's home directory
func ExpandHomeDir(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if path == "~" {
		return homeDir, nil
	}

	return filepath.Join(homeDir, path[2:]), nil
}

// NormalizePath normalizes a path to use the OS-specific separator
func NormalizePath(path string) string {
	return filepath.Clean(path)
}

// GetDefaultDownloadDir returns the default download directory for the OS
func GetDefaultDownloadDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	var downloadDir string
	switch runtime.GOOS {
	case "windows":
		downloadDir = filepath.Join(homeDir, "Downloads")
	case "darwin":
		downloadDir = filepath.Join(homeDir, "Downloads")
	case "linux":
		// Try XDG_DOWNLOAD_DIR first
		xdgDownload := os.Getenv("XDG_DOWNLOAD_DIR")
		if xdgDownload != "" {
			downloadDir = xdgDownload
		} else {
			downloadDir = filepath.Join(homeDir, "Downloads")
		}
	default:
		downloadDir = filepath.Join(homeDir, "Downloads")
	}

	return downloadDir, nil
}

// EnsureDir ensures a directory exists, creating it if necessary
func EnsureDir(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(path, 0755)
		}
		return err
	}

	if !info.IsDir() {
		return errors.New("path exists but is not a directory")
	}

	return nil
}

// IsWritable checks if a directory is writable
func IsWritable(path string) bool {
	// Try to create a temporary file
	testFile := filepath.Join(path, ".write_test")
	file, err := os.Create(testFile)
	if err != nil {
		return false
	}
	file.Close()
	os.Remove(testFile)
	return true
}

// GetDiskSpace returns available disk space in bytes
func GetDiskSpace(path string) (uint64, error) {
	// This is a simplified version
	// For production, you'd want to use syscall for accurate disk space
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		path = filepath.Dir(path)
	}

	// Platform-specific disk space checking would go here
	// For now, return a large number as placeholder
	return 100 * 1024 * 1024 * 1024, nil // 100 GB
}

// ListDirectories returns a list of directories in the given path
func ListDirectories(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}

	return dirs, nil
}

// GetParentDir returns the parent directory of the given path
func GetParentDir(path string) string {
	return filepath.Dir(path)
}

// JoinPath joins path elements and normalizes the result
func JoinPath(elements ...string) string {
	return filepath.Join(elements...)
}

// IsHidden checks if a file or directory is hidden
func IsHidden(name string) bool {
	if runtime.GOOS == "windows" {
		// On Windows, check file attributes (simplified)
		return strings.HasPrefix(name, ".")
	}
	// On Unix-like systems, hidden files start with a dot
	return strings.HasPrefix(name, ".")
}
