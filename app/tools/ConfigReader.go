package tools

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigReader struct{}

func (c *ConfigReader) ReadFileConfiguration(pathFile string, i interface{}) error {
	rootPath, err := os.Getwd()
	if err != nil {
		return err
	}

	finalPath := rootPath + pathFile

	file, err := os.Open(finalPath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&i)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
