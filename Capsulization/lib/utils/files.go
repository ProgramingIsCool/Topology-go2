package utils

// not in this  packager

import (
	"io"
	"os"
)

// LoadFile loads the entire file and returns its content as a byte slice ([]byte).
// This function has a single responsibility: read file data from disk into memory.
func FileContent(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
