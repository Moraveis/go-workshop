// Exercise 1.06: Declaring Multiple Variables from a Function

package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {

	return false, "info", time.Now()
}

func main() {

	// Type only >> all didentifiers must have the same time
	var start, middle, end float32
	fmt.Println(start, middle, end)

	// Initial values mixed types
	var name, left, right, top, bottom = "one", 1, 1.5, 2, 2.5
	fmt.Println(name, left, right, top, bottom)

	// works with functions also
	var Debug, LogLevel, startUpTime = getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)

}