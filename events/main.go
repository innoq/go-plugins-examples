package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/innoq/go-plugins-examples/utils"
	"github.com/robertkrimen/otto"
)

var listeners map[string][]string
var scripts map[string]string

func init() {
	scripts = make(map[string]string)
	listeners = make(map[string][]string)
	for _, event := range []string{"create", "read", "update", "delete"} {
		listeners[event] = make([]string, 0)
	}
	err := loadPlugins()
	if err != nil {
		panic(err)
	}
}

func main() {
	vm := otto.New()
	vm.Set("save", saveMethod)
	generator := utils.NewDataGenerator()
	for i := 0; i < 42; i++ {
		fmt.Println("")
		data := generator.Next()
		vm.Set("data", data)
		_, err := vm.Run(scripts["creator"])
		if err != nil {
			panic(err)
		}
		fmt.Println("")
	}
}

func read(vm *otto.Otto, script string) {
	vm.Run(script)
}

func saveMethod(data map[string]string) error {
	log.Printf("saving some data %s", data)
	return notifyListener(data)
}

func notifyListener(data map[string]string) error {
	event := data["event"]
	if listenerScripts, ok := listeners[event]; ok {
		vm := otto.New()
		for _, name := range listenerScripts {
			log.Printf("notify %s about %s event", name, event)
			err := vm.Set("data", data)
			if err != nil {
				return err
			}

			_, err = vm.Run(scripts[name])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func loadPlugins() error {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			info, err := utils.ReadJSON(path.Join(file.Name(), "info.json"))
			if err != nil {
				return err
			}

			script, err := utils.ReadFile(path.Join(file.Name(), "script.js"))
			if err != nil {
				return err
			}
			scripts[file.Name()] = script
			events := info["events"].([]interface{})
			for _, eventEntry := range events {
				event := eventEntry.(string)
				listeners[event] = append(listeners[event], file.Name())
				log.Printf("register %s for event %s", file.Name(), event)
			}
		}
	}
	return nil
}
