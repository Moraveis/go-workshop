// Exercise 2.06: switch Statements and Multiple case Values

package main

import (
	"fmt"
	"time"
)

func main() {

	dayBorn := time.Sunday

	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Born on a weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("Born on a weekened")
	default:
		fmt.Println("Error, day born not valid")
	}

}
