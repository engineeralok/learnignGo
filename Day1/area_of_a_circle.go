package main

import (
	"fmt"
	"math"
)

func main() {
	var area, radius float32
	// const pi = 3.14
	fmt.Println("Input radius")
	fmt.Scanf("%f", &radius)
	area = math.Pi * radius * radius
	fmt.Printf("radius of %f meters; area is %f square meters\n", radius, area)
	fmt.Print("radius of ", radius, " meters; area is ", area, " square meters\n")
	fmt.Println("radius of ", radius, "meters; area is ", area, "square meters")
}
