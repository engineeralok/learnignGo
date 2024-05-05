package main

import "fmt"

func main() {
	var p, q, r bool // default initialized to false
	var a, b, c int = 0, 1, 2
	fmt.Println("Bool Logic\n p q r\n________________________")

	fmt.Println(p, q, r, "initialized values")

	p = true
	q = true
	r = p && q
	fmt.Println(p, q, r, "p && q")

	p = !p //negation
	r = p && q

	fmt.Println(p, q, r, "p && q")

	r = p || q

	fmt.Println(p, q, r, "p || q")

	fmt.Println("a = ", a, " b = ", b, " c = ", c)
	fmt.Println("a == b ", a == b, " a < c ", a < c, " c != b ", c != b)
}
