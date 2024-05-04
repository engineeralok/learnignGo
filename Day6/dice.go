// implentation of a simulated dice roll
// Will use random number generator to simulate dice rolls

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// fmt.Println(rand.Intn(6) + 1)

	var pair, howMany int

	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10000000; i++ {
		pair = rand.Intn(6) + rand.Intn(6) + 2
		if pair == 7 {
			howMany += 1
		}
	}

	fmt.Println(howMany)
	fmt.Printf("Probability of rolling a seven is %4g \n", float64(howMany)/10000000.0)
	// %4g is a width of 4 and a precision of g which is a general format that prints a number in a way that is easy to read.

}


