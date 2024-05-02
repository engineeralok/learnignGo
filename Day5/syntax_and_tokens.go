package main

import (
	"fmt"
)

var myInt int = 5 // camel case unlike my_int which is snake case // global variable
func main() {

	fmt.Println("Will try some simple ideas ")
	fmt.Println("Print the global", myInt)
	{
		var myInt int = 6 //redclared outer declaration hidden
		fmt.Println("Print the inner", myInt)
	}
	fmt.Println("Print the global", myInt)

}

/*
	Go has 25 keywords: break, default, func, interface,select,
	package, switch, const, fallthrough, if, range, type,
	continue, for, import, return, var

	These words cannot me use as identifiers.
*/
