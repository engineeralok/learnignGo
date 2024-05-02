// Exploring basic types and declarations
package main

import "fmt"

func main() {
	fmt.Println("Will use basic types and declarations")
	var myInt int    //simple declaration of int
	anotherInt := 36 //implicit declaration of int
	myFloat := 8.95
	var data1, data2, data3 float32 = 0.1, 1.5, 3.5
	var mData float64
	var unsignedNumber uint
	fmt.Println("Print myInt", myInt)
	fmt.Println("Print anotherInt", anotherInt)
	fmt.Println("Print myFloat", myFloat)
	fmt.Println("Print data", data1, data2, data3)
	fmt.Println("Print mData", mData)
	fmt.Println("Print unsignedNumber", unsignedNumber)
	myInt = myInt + anotherInt
	fmt.Println("Print myInt", myInt)
	unsignedNumber = uint(myInt) // need conversion in GO
	fmt.Println("Print unsignedNumber", unsignedNumber)
	unsignedNumber = uint(-myInt) // a very large number: why???
	fmt.Println("Print unsignedNumber", unsignedNumber)
	mData = float64(data1 + data2 + data3)
	fmt.Println("Print mData", mData)

}

/*
	How to declare variables in some basic types

	Types:
	any bool byte comparable
	complex64 complex128
	float32 float64
	int int8 int16 int32 int64
	rune
	uint uint8 uint16 uint32 uint64 uintptr
	string

*/
