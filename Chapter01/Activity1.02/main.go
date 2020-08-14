/**

Activity 1.02: Pointer Value Swap

1. Call the swap function, ensuring you are passing a pointer.
2. In the swap function, assign the values to the other pointer, ensuring you dereference the values.
The following is the expected output:

AC :
	a == 10 and b == 5
	>> true and true
*/

package main

import "fmt"

func main() {

	a, b := 5, 10

	swap(&a, &b)
	fmt.Println(a == 10, b == 5)

}

func swap(a *int, b *int) {
	// swap the values here

	aux := *a
	*a = *b
	*b = aux

}
