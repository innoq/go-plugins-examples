package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("Hello World!")
}

// Greeter - exported
var Greeter greeting
