package main

import (
	"fmt"
)

func main() {
	var miles, yards int
	// fmt.Println("Enter Miles: ")
	// fmt.Scan(&miles)
	// fmt.Println("Enter Yards: ")
	// fmt.Scan(&yards)

	// kilometers := convertToKm(miles, yards)

	// fmt.Println("Distance in Kilometers: ", kilometers)

	for i := 0; i < 10; i++ {
		fmt.Println("Converts Miles and Yards to Kilometers : input 2 numbers")
		fmt.Println("If Miles is negative program will exit")
		fmt.Println("Enter Miles: ")
		fmt.Scan(&miles)
		fmt.Println("Enter Yards: ")
		fmt.Scan(&yards)
		if miles < 0 || yards < 0 {
			fmt.Println("Exiting")
			return
		}
		fmt.Println("Distance in Kilometers: ", convertToKm(miles, yards))
	}
}

// convertToKm converts miles and yards to kilometers and returns the result.
// Parameters:
// - miles: the number of miles to convert
// - yards: the number of yards to convert
// Return:
// - kiloMeter: the converted distance in kilometers

// We need to cast the result of the multiplication to float64 because the
// result of the division is a float64.
func convertToKm(miles int, yards int) float64 {
	kiloMeter := (1.609 * (float64(miles) + float64(yards)/1760.0))
	return kiloMeter
}
