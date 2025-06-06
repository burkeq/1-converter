package files

import (
	"os"
	"strings"
)

func ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func CheckJsonExtension(filename string) bool {
	if !strings.HasSuffix(filename, ".json") {
		return false
	}
	return true
}
