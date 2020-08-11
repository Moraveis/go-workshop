// Exercise 1.02: Declaring a Variable Using var

package main

import (
	"fmt"
)

var foo string = "bar"

func main() {

	var baz string = "qux"

	fmt.Println(foo, baz)
}
