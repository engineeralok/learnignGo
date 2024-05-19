package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var move, machineMove int
	const (
		rock     = 0
		paper    = 1
		scissors = 2
	)
	const (
		cRock     = "R"
		cPaper    = "P"
		cScissors = "S"
	)

	var cMove string
	var drows, wins, machineWins int
	var rounds int

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many rounds would you like to play? ")
	text, _ := reader.ReadString('\n')
	fmt.Sscanf(text, "%d", &rounds)

	for i := 0; i < rounds; i++ {
		// User Plays
		fmt.Println("\n Round", i+1, ": Choose either R, P, or S")

		text, _ = reader.ReadString('\n')
		cMove = text[0:1]
		fmt.Scanf("%c\n", &cMove)

		if cMove == cRock {
			move = rock
		} else if cMove == cPaper {
			move = paper
		} else if cMove == cScissors {
			move = scissors
		} else {
			fmt.Println("--> Invalid move")
			i--
			continue
		}

		// Machine Plays
		rand.Seed(time.Now().UnixNano())
		machineMove = rand.Intn(3)

		if machineMove == 0 {
			cMove = cRock
		} else if machineMove == 1 {
			cMove = cPaper
		} else if machineMove == 2 {
			cMove = cScissors
		}
		fmt.Printf("Machine plays %s\n", cMove)

		// Determine result
		if move == machineMove {
			fmt.Println("Draw")
			drows++
		} else if (move == paper) && (machineMove == scissors) {
			machineWins++
			fmt.Println("Machine wins")
			continue
		} else if (move == scissors) && (machineMove == rock) {
			machineWins++
			fmt.Println("Machine wins")
			continue
		} else if (move == rock) && (machineMove == paper) {
			machineWins++
			fmt.Println("Machine wins")
			continue
		} else if (move == rock) && (machineMove == scissors) {
			wins++
			fmt.Println("You win")
			continue
		} else if (move == paper) && (machineMove == rock) {
			wins++
			fmt.Println("You win")
			continue
		} else if (move == scissors) && (machineMove == paper) {
			wins++
			fmt.Println("You win")
			continue
		} else {
			wins++
			fmt.Println("You win")
		}
	}
	fmt.Println("\nAfter ", rounds, " rounds, you won", wins, " times, the machine won ", machineWins, " times, and there were ", drows, " draws")

}
