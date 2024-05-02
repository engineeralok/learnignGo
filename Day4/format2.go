// Exploring use of formatting floating point
// %e - scientific notation
// %f - decimal float
// %.nf - n is int precision
// %m.nf - width m precision n
// %g - concise floating point
// 10.2g - width 10 precision 2

// %s - string
// %ns - width right justify n string
// %ns - width left justify n string
// %c - character
// %t - boolean
// %U - unicode
// %#U - unicode with character

package main

import (
	"fmt"
)

func main() {
	x := 345.6789
	fmt.Println("Here is Println with defaults x = ", x)
	fmt.Printf("Here is Printf with f (decimal float) format: x = %f\n", x)
	fmt.Printf("Here is Printf with e (scientific notation) format: x = %e\n", x)
	fmt.Printf("Here is Printf with g (concise floating point) format: x = %g\n", x)

	x = 1000 * x
	fmt.Printf("Here is Printf with g (concise floating point) format: x = %g\n", x)
	fmt.Printf("Here is Printf with 10.2g (concise floating point) format: x = %10.2g\n", x)

}
