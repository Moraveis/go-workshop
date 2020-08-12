// Exercise 1.08: Changing Multiple Values at Once

package main

import "fmt"

func main() {

	query, limit, offset := "bat", 10, 0
	fmt.Println(query, limit, offset)

	query, limit, offset = "ball", offset, 20
	fmt.Println(query, limit, offset)
}
