// Exercise 1.13: Getting a Pointer

package main

import (
	"fmt"
	"time"
)

func main() {

	// Create a pointer using `*`
	var count1 *int

	// Create a pointer using the func `new(type)`
	count2 := new(int)

	// Get pointer from existing variable using `&`
	countTemp := 5
	count3 := &countTemp

	// Get pointer without tem variable
	t := &time.Time{}

	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time: %#v\n", t)

}
