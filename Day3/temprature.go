package main

import (
	"fmt"
)

func main() {
	var from, to, by int32 = 0, 250, 10
	var centigrade, fahrenheit float32
	fahrenheit = float32(from)

	for fahrenheit <= float32(to) {
		centigrade = float32(fahrenheit-32) * 5 / 9
		fmt.Printf("%g°F = %g°C\n", fahrenheit, centigrade)
		fahrenheit += float32(by)
	}
}
