package main

import (
	"github.com/innoq/go-plugins-examples/utils"
	"github.com/robertkrimen/otto"
)

func main() {

	vm := otto.New()
	observerScript, err := utils.ReadFile("./observer.js")
	if err != nil {
		panic(err)
	}

	vm.Set("save", createSaveMethod(observerScript))
	creatorScript, err := utils.ReadFile("./creator.js")
	if err != nil {
		panic(err)
	}

	generator := utils.NewDataGenerator()
	for i := 0; i < 10; i++ {
		data := generator.Next()
		vm.Set("data", data)
		_, err := vm.Run(creatorScript)
		if err != nil {
			panic(err)
		}
	}
}

func read(vm *otto.Otto, script string) {
	vm.Run(script)
}

func createSaveMethod(script string) func(map[string]string) error {
	vm := otto.New()
	return func(data map[string]string) error {
		event := make(map[string]string)
		event["updated_type"] = data["type"]
		event["updated_id"] = data["id"]
		vm.Set("event", event)
		_, err := vm.Run(script)
		if err != nil {
			return err
		}
		return nil
	}
}
