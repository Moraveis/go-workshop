// Exercise 1.04: Skipping the Type or Value When Declaring Variables

package main

import (
	"fmt"
	"time"
)

var (
	Debug       bool
	LogLevel    = "info"
	startUpTime = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}
