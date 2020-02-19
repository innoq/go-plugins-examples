package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/robertkrimen/otto"
)

// ReadJSON - reads the content of a JSON file, returns it as a map
func ReadJSON(path string) (map[string]interface{}, error) {
	content, err := ReadFile(path)
	if err != nil {
		return make(map[string]interface{}), err
	}
	var jsonContent map[string]interface{}
	err = json.Unmarshal([]byte(content), &jsonContent)
	if err != nil {
		return make(map[string]interface{}), err
	}
	return jsonContent, nil
}

// ReadFile - reads the content of a file, returns it as a string
func ReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// InjectEnvironmentVariables - inject all requeste Environment Variables
func InjectEnvironmentVariables(envKeys []interface{}, vm *otto.Otto) error {
	envs := make(map[string]string)
	for _, element := range envKeys {
		k := element.(string)
		v := os.Getenv(k)
		envs[k] = v
	}
	err := vm.Set("env", func(k string) string {
		if v, ok := envs[k]; ok {
			return v
		}
		return ""
	})
	if err != nil {
		return err
	}
	return nil
}
