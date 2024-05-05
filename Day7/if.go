/*
	if and if-else are two flow of control statements that allow the program to be
	breaching instead of linear.
	or
	breaching the flow of control if a certain condition is met

	They also involve controlling a block

	some difference from C

	C cannot have a simple statement before the if expression.
	C needs (...exp) and does not always need a {...} block.


*/

package main

import "fmt"

func main() {
	fmt.Println("Results of if ")
	x := 5
	y := 6
	i := 3 // test scope

	fmt.Println("x is ", x, " y is ", y, " i is ", i)

	if x < y {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("x is not less than y")
	}

	if i := 9; i < x { // i here hides outer scope i
		fmt.Println("i is less than x", "i is ", i, "x is ", x)
	} else {
		fmt.Println("i is not less than x", "i is ", i, "x is ", x)
	}

	fmt.Println("i is ", i) // i here will print 3

}
