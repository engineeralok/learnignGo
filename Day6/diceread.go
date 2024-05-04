package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var pair, vPair, howMany int
	rand.Seed(time.Now().UnixNano())
	fmt.Println("readin pair of dice value (2-12):")
	fmt.Scanf("%d", &vPair)
	for i := 0; i < 10000000; i++ {
		pair = rand.Intn(6) + rand.Intn(6) + 2
		if pair == vPair {
			howMany++
		}
	}
	fmt.Println(howMany)
	fmt.Printf("Probability of rolling a %d is %4g \n", vPair, float64(howMany)/10000000.0)
}
