package main

import (
	"fmt"
	"sort"

	"github.com/robertkrimen/otto"
)

func main() {
	// Create a new otto interpreter instance.
	vm := otto.New()

	// This is where the magic happens!
	vm.SetDebuggerHandler(func(o *otto.Otto) {
		// The `Context` function is another hidden gem - I'll talk about that in
		// another post.
		c := o.Context()

		// Here, we go through all the symbols in scope, adding their names to a
		// list.
		var a []string
		for k := range c.Symbols {
			a = append(a, k)
		}

		sort.Strings(a)

		// Print out the symbols in scope.
		fmt.Printf("symbols in scope: %v\n", a)
	})

	// Here's our script - very simple.
	s := `
    var a = 1;
    var b = 2;
    debugger;
  `

	// When we run this, we should see all the symbols printed out (including
	// `a` and `b`).
	if _, err := vm.Run(s); err != nil {
		panic(err)
	}
}
