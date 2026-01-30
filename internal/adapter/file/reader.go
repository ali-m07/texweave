package file

import (
	"os"
)

// ReadAll reads the entire file at path.
func ReadAll(path string) ([]byte, error) {
	return os.ReadFile(path)
}
