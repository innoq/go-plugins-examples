package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"

	"github.com/innoq/go-otto-examples/utils"
	"github.com/robertkrimen/otto"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {

	var err error

	err = runPlugin("twitter-plugin")
	if err != nil {
		return err
	}

	err = runPlugin("basic-auth-plugin")
	if err != nil {
		return err
	}

	return nil
}

func runPlugin(plugin string) error {

	fmt.Printf("\n\nrunning Plugin %s: \n\n", plugin)
	info, err := utils.ReadJSON(path.Join(plugin, "info.json"))
	if err != nil {
		return err
	}
	log.Printf("info: %s", info)
	fmt.Println("")
	vm, err := getVM(info)
	if err != nil {
		return err
	}

	script, err := utils.ReadFile(path.Join(plugin, "client.js"))
	if err != nil {
		return err
	}

	_, err = vm.Run(script)
	if err != nil {
		return err
	}

	return nil
}

func getVM(info map[string]interface{}) (*otto.Otto, error) {
	vm := otto.New()

	if envKeyValue, ok := info["env_variables"]; ok {
		envKeys := envKeyValue.([]interface{})
		err := utils.InjectEnvironmentVariables(envKeys, vm)
		if err != nil {
			return nil, err
		}
	}

	err := vm.Set("LOG", func(message string) {
		log.Println(message)
	})
	if err != nil {
		return nil, err
	}

	err = vm.Set("POST", loadPostRequest(info))
	if err != nil {
		return nil, err
	}

	err = vm.Set("GET", loadGetRequest(info))
	if err != nil {
		return nil, err
	}

	return vm, nil
}

// Method - type of a http Method
type Method func(request map[string]interface{}) map[string]interface{}

func loadPostRequest(info map[string]interface{}) Method {
	return func(request map[string]interface{}) map[string]interface{} {
		response := make(map[string]interface{})
		whitelist := info["whitelist"].([]interface{})
		requestURL := request["host"].(string)
		err := checkWhitelist(requestURL, whitelist)
		if err != nil {
			response["error"] = err.Error()
			return response
		}

		client, err := utils.NewClient(request)
		if err != nil {
			response["error"] = err.Error()
			return response
		}
		content := ""
		if request["content"] != nil {
			content = request["content"].(string)
		}

		contentType := "application/json"
		if request["contentType"] != nil {
			contentType = request["contentType"].(string)
		}

		body, err := client.Post(content, contentType)
		if err != nil {
			response["error"] = err.Error()
			return response
		}

		response["body"] = body
		return response
	}
}

func loadGetRequest(info map[string]interface{}) Method {
	return func(request map[string]interface{}) map[string]interface{} {
		response := make(map[string]interface{})
		whitelist := info["whitelist"].([]interface{})
		requestURL := request["host"].(string)
		err := checkWhitelist(requestURL, whitelist)
		if err != nil {
			response["error"] = err.Error()
			return response
		}

		client, err := utils.NewClient(request)
		if err != nil {
			response["error"] = err.Error()
			return response
		}

		body, err := client.Get()
		if err != nil {
			response["error"] = err.Error()
			return response
		}

		response["body"] = body
		return response
	}
}

func checkWhitelist(requestURL string, whitelist []interface{}) error {
	url, err := url.Parse(requestURL)
	if err != nil {
		return err
	}
	host := fmt.Sprintf("%s://%s", url.Scheme, url.Host)
	for _, entry := range whitelist {
		if entry.(string) == host {
			return nil
		}
	}
	return fmt.Errorf("accessing %s is blocked", host)
}
