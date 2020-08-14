// Exercise 1.14: Getting a Value from a Pointer

package main

import (
	"fmt"
	"time"
)

func main() {

	// Declaring pointers
	var count1 *int

	count2 := new(int)

	countTemp := 5
	count3 := &countTemp

	t := &time.Time{}

	// Checks for `nil`
	if count1 != nil {
		fmt.Printf("count1 (value): %#v\n", *count1)
	}

	if count2 != nil {
		fmt.Printf("count2 (value): %#v\n", *count2)
	}

	if count3 != nil {
		fmt.Printf("count3 (value): %#v\n", *count3)
	}

	if t != nil {
		fmt.Printf("time (value): %#v\n", *t)

		fmt.Printf("time (value): %#v\n", t.String())
	}

}
