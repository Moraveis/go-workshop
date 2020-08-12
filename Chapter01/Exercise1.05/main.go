//Exercise 1.05: Implementing Short Variable Declaration

package main

import (
	"fmt"
	"time"
)

func main() {

	Debug := false
	LogLevel := "info"
	startUpTime := time.Now()

	fmt.Println(Debug, LogLevel, startUpTime)
}
