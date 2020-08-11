// Exercise 1.01: Using Variables, Packages, and Functions to Print Stars

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(5) + 1

	stars := strings.Repeat("*", r)

	fmt.Print(stars)

}
