package main

import (
	"log"

	"github.com/robertkrimen/otto"
)

func main() {
	vm := otto.New()

	err := vm.Set("log", logJS)
	if err != nil {
		panic(err)
	}

	vm.Run(`
		console.log("logging with JS!");
		log("logging with Golang!");
	`)

	jsDate, err := vm.Run(`
	(function(){
		date = new Date();
		return date;
	})();
	`)
	if err != nil {
		panic(err)
	}
	log.Printf("jsDate: %s", jsDate)

	dataMap := make(map[string]interface{})
	dataMap["foo"] = "bar"
	dataMap["one"] = "1"
	dataMap["two"] = "2"

	err = vm.Set("dataMap", dataMap)
	if err != nil {
		panic(err)
	}

	value, err := vm.Run(`
	(function(){
		var keys = [];
		for(k in dataMap) {
			log(k + ": " + dataMap[k]);
			keys.push(k);
		}
		return keys;
	})();
	`)

	keys, err := value.Export()
	if err != nil {
		panic(err)
	}

	keyArray := keys.([]string)
	log.Printf("keys: %s", keyArray)

}

func logJS(content string) {
	log.Println(content)
}
