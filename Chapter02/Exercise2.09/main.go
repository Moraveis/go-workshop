// Exercise 2.09: Looping Over Arrays and Slices

package main

import "fmt"

func main() {

	names := []string{"Jim", "Jane", "Joe", "June"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
