// Write a GO program that improves on the class program for Rock-Paper-Scissors.

// This program should adjust its strategy to that of its opponent. For example, if it sees that the user is only playing rock then it should switch to playing scissors. This might be considered a simple form of machine learning. One way to do this is to keep track of the users' plays and after some time, say every ten plays, adjust the machine strategy to play using this knowledge. 

// Assume a match is 100 plays.

package main

import(
	
	"fmt"
	"math/rand"
	"time"
)

