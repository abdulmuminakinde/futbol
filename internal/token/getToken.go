package token

import (
	"os"
	"strings"
)

func GetToken() string {
	file := GetFilePath()
	data, err := os.ReadFile(file)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data))
}
