// Type Inference Gone Wrong

package main

import "math/rand"

func main() {

	// var seed = 1234456789 // cannot use seed (type int) as type int64 in argument to rand.Seedgo
	var seed int64 = 1234456789

	rand.Seed(seed)

}
