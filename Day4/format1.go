//%d - use decimal base 10
//%b - use binary base 2
//%o - use octal base 8
//%O - use octal base 8
//%x - use hex base 16
//%f - use float
//%s - use string
//%c - use character
//%t - use boolean
//%U - use unicode
//%#U - use unicode with character

package main

import (
	"fmt"
)

func main() {
	i := 55
	fmt.Println("Here is Println with defaults i = ", i)
	fmt.Printf("Here is Printf with decimal format: i = %d\n", i)
	fmt.Printf("Here is Printf with width 10d format: i = %10d\n", i)
	fmt.Printf("Here is Printf with binary format: i = %b\n", i)
	fmt.Printf("Here is Printf with o octal format: i = %o\n", i)
	fmt.Printf("Here is Printf with O octal format: i = %O\n", i)
	fmt.Printf("Here is Printf with hex format: i = %x\n", i)
	fmt.Printf("Here is Printf with character format: i = %c\n", i)
}
