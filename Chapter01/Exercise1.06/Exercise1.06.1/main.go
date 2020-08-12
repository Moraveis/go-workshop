// Exercise 1.06: Declaring Multiple Variables with a Short Variable Declaration

package main

import (
	"fmt"
	"time"
)

func main() {
	Debug, LogLevel, startUpTime := false, "info", time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)
}
