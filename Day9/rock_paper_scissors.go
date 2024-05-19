// rune - rune is a datatype that represents a character
// rune is an alias for int32

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var move rune
	var machineMove int
	fmt.Println("Choose either R P or S")
	fmt.Scanf("%c", &move)
	if move == 'R' || move == 'r' {
		fmt.Printf("my move is %c\n", move)
	} else if move == 'P' || move == 'p' {
		fmt.Printf("my move is %c\n", move)
	} else if move == 'S' || move == 's' {
		fmt.Printf("my move is %c\n", move)
	} else {
		fmt.Println("invalid move")
	}
	rand.Seed(time.Now().UnixNano())
	machineMove = rand.Intn(3)
	if machineMove == 0 {
		fmt.Println("machine move is R")
	} else if machineMove == 1 {
		fmt.Println("machine move is P")
	} else if machineMove == 2 {
		fmt.Println("machine move is S")
	}

	if machineMove == 0 && move == 'S' || machineMove == 0 && move == 's' {
		fmt.Println("machine wins")
	} else if machineMove == 1 && move == 'R' || machineMove == 1 && move == 'r' {
		fmt.Println("machine wins")
	} else if machineMove == 2 && move == 'P' || machineMove == 2 && move == 'p' {
		fmt.Println("machine wins")
	} else if machineMove == 0 && move == 'P' || machineMove == 0 && move == 'p' {
		fmt.Println("you win")
	} else if machineMove == 1 && move == 'S' || machineMove == 1 && move == 's' {
		fmt.Println("you win")
	} else if machineMove == 2 && move == 'R' || machineMove == 2 && move == 'r' {
		fmt.Println("you win")
	} else {
		fmt.Println("draw")
	}

}
