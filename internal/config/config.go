package config

import (
	"encoding/json"
	"os"
	"strings"
)

type Interface interface {
	GetStringData(key string) (value string)
	GetIntData(key string) (value int)
	GetBooleanData(key string) (value bool)
}

type Config struct {
	configData map[string]interface{}
}

func (c *Config) get(key string) (value interface{}) {
	data := c.configData
	for _, key := range strings.Split(key, ":") {
		value = data[key]
	}
	return
}

func (c *Config) GetStringData(key string) string {
	value := c.get(key)
	return value.(string)
}

func (c *Config) GetIntData(key string) int {
	value := c.get(key)
	return value.(int)
}

func (c *Config) GetBooleanData(key string) bool {
	value := c.get(key)
	return value.(bool)
}

func LoadFile(fileName string) (config Interface, err error) {
	var fileData []byte
	fileData, err = os.ReadFile(fileName)
	if err == nil {
		data := json.NewDecoder(strings.NewReader(string(fileData)))
		keyValue := map[string]interface{}{}
		err = data.Decode(&keyValue)
		if err == nil {
			config = &Config{configData: keyValue}
		}
	}
	return
}
