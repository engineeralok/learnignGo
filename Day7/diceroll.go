package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var pair, vPair, howmany int

	fmt.Println("What you want to get? between 1 and 12 : ")
	fmt.Scanf("%d", &vPair)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 1000000; i++ {
		pair = rand.Intn(6) + rand.Intn(6) + 2
		if pair == vPair {
			howmany++
		}

	}

	fmt.Println("The probability of getting", vPair, "when rolling dice 1000000 times is", (float64(howmany) / 1000000))

}
