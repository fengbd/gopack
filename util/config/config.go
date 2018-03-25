package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Parse a file to config
func Parse(path string, out interface{}) error {
	if !IsExist(path) {
		return os.ErrNotExist
	}

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, out)
}

// Save config to file
func Save(path string, in interface{}) error {
	data, err := yaml.Marshal(in)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, os.ModePerm)
}

// IsExist is returns true if path exist
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
