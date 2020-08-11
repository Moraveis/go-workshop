// use `go build` to compile, and `./example01>` to run

package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var helloList = []string{
	"Hello, world",
	"Καλημέρα κόσμε",
	"こんにちは世界",
	"سلام دنیا‎",
	"Привет, мир",
}

func main() {
	// Seed random number generator using the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number in the range of out list
	index := rand.Intn(len(helloList))

	// Call a function and receive multiple return values
	msg, err := hello(index)

	// Handle any errors
	if err != nil {
		log.Fatal(err)
	}

	// Print our message to the console
	fmt.Println(msg)

}

func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {

		// Create an error, convert the int type to a string
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}

	return helloList[index], nil
}
