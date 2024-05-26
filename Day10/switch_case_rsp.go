// The Switch is a multi-way branching statement.

// Switch statements:
// SwitchStmt = ExprSwitchStmt | TypeSwitchStm

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
	var draws, wins, machineWins int
	var rounds int

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How many rounds do you want to play? ")
	fmt.Scanf("%d", &rounds)

	for i := 0; i < rounds; i++ {

		fmt.Println("\nRound ", i+1, " : Choose either R, P or S")

		fmt.Scanf("%s", &cMove)

		reader.ReadString('\n')

		fmt.Printf("You played %s\n", cMove)

		if cMove == cRock || cMove == cPaper || cMove == cScissors {
			move = -1
			switch cMove {
			case cRock:
				move = rock
			case cPaper:
				move = paper
			case cScissors:
				move = scissors
			}
		} else {
			fmt.Println("--> Invalid move")
			i--
			continue
		}

		//Machine plays
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

		switch move {
		case rock:
			if machineMove == paper {
				fmt.Println("Machine wins")
				machineWins++
			} else if machineMove == scissors {
				fmt.Println("User wins")
				wins++
			} else {
				fmt.Println("Draw")
				draws++
			}
		case paper:
			if machineMove == scissors {
				fmt.Println("Machine wins")
				machineWins++
			} else if machineMove == rock {
				fmt.Println("User wins")
				wins++
			} else {
				fmt.Println("Draw")
				draws++
			}
		case scissors:
			if machineMove == rock {
				fmt.Println("Machine wins")
				machineWins++
			} else if machineMove == paper {
				fmt.Println("User wins")
				wins++
			} else {
				fmt.Println("Draw")
				draws++
			}
		}
	}

	fmt.Println("\nAfter ", rounds, " rounds: \n", "you win: ", wins, "\n", "machine wins: ", machineWins, "\n", "draws: ", draws)
}
