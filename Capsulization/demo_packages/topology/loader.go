package topology

import (
	"fmt"
	"io"
	"os"
)

// LoadFile loads the entire file and returns its content as a byte slice ([]byte).
// This function has a single responsibility: read file data from disk into memory.
func LoadFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}

	return data, nil
}
