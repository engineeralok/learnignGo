// Will use random number generator to simulate dice rolls
// Add time to seed the random number generator
// This will allow the program to produce different results each run

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println(time.Now().UnixNano())

	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// The above computes the time in nanoseconds (1 billionth of a second) since the start of the unix epoch (Jan1, 1970).
	// This will allow the program to produce different results each run
	// basically each time the program is run it will be a different seed for the random number generator

	for i := 0; i < 10; i++ {
		fmt.Printf("%d \t", rand.Intn(6)+1)
	}

	fmt.Println("BYE")
}
