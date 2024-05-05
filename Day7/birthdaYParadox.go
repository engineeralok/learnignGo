/*

Write a GO program that uses random number generation to solve a classical probability problem.

Assume a room with 20 people. What is the probability that two of them have the same birthday?
For simplicity, say a year has 365 days.
Day 1 is January 1 and day 365 is December 31.

So, a birthday will be a number from 1 to 365.  Now compute whether 20 people at random will have the same birthday.
Do this 10000 times to get a reasonable sample.
Also do this for 10, 20, 30 ,.. 100 people in the room and print this as a table.  So the output should be:

People in the room       Probability of 2 (or more) having the same birthday

10                                          0.xyz

20                                          0.xyz

…

A typical bar bet is made on where the probability is greater than one-half.

*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	fmt.Println("People in the room\tProbability of 2 (or more) having the same birthday")

	// Run the simulation for 10 to 100 people in steps of 10
	for people := 10; people <= 100; people += 10 {
		simulations := 10000
		successCount := 0

		for i := 0; i < simulations; i++ {
			if hasDuplicateBirthday(people) {
				successCount++
			}
		}

		// Calculate probability
		probability := float64(successCount) / float64(simulations)
		// birthdaysFound := successCount
		fmt.Printf("%d\t\t\t\t%.3f \t\t\t\t%d\n", people, probability,)
	}
}

// hasDuplicateBirthday simulates assigning random birthdays to 'people' number of people
// and checks if at least two have the same birthday.
func hasDuplicateBirthday(people int) bool {
	birthdays := make(map[int]bool)
	for i := 0; i < people; i++ {
		birthday := rand.Intn(365) + 1 // Random birthday between 1 and 365
		if birthdays[birthday] {
			return true
		}
		birthdays[birthday] = true
	}
	return false
}
