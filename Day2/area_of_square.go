package main

import (
	"fmt"
)

func main() {
	var bhuja, area float32
	fmt.Print("bhuja ki value likho : ")
	fmt.Scanf("%f", &bhuja)
	area = bhuja * bhuja
	fmt.Println("bhuja ki value : ", bhuja, " meter. Area of square : ", area, " meter square")
	fmt.Printf("bhuja ki value : %g meter. Area of square : %g meter square", bhuja, area)
}
