// Activity 1.03: Message Bug

package main

import "fmt"

func main() {
	count := 4

	var message string

	if count > 5 {
		message = "Greater than 5"
	} else {
		message = "Not greater than 5"
	}

	fmt.Println(message)
}
