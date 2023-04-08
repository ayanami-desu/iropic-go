package utils

import (
	json "github.com/json-iterator/go"
	"os"
)

var Json = json.ConfigCompatibleWithStandardLibrary

// WriteJsonToFile write struct to json file
func WriteJsonToFile(dst string, data interface{}) bool {
	str, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return false
	}
	err = os.WriteFile(dst, str, 0777)
	if err != nil {
		return false
	}
	return true
}
