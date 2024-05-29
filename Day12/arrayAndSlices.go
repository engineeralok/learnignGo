/*
	An array is a homogeneous data structure that lets you store elements of the same type.

	Arrays are used to store collections of data.

	Arrays are used to store a fixed number of values of the same type.

	Arrays are zero-indexed.

	Arrays are declared using the keyword array.

	Arrays are created using the following syntax:

		var identifier [length]type

		For example:
		var arr [5]int
		var arr1 = [5]int{1, 2, 3, 4, 5}
		var arr2 = [...]int{1, 2, 3, 4, 5}
		var arr3 = [...]int{1, 2, 3, 4, 5: 10}
		var arr4 = [...]int{1, 2, 3, 4, 5: 10, 11}
		var arr5 = [...]int{1, 2, 3, 4, 5: 10, 11, 12, 13, 14, 15}
		var arr6 = [...]int{1, 2, 3, 4, 5: 10, 11, 12, 13, 14, 15: 20}
		var arr7 = [...]int{1: 10}
		var arr8 = [...]int{1: 10, 2: 20}
		var arr9 = [...]int{1: 10, 2: 20, 3: 30}

		An example in Go is var data [5] float 64 which creates: data[0], data[1],
		data[2], data[3], and data[4] all initialized to 0.0 They can be accessed with an indes such as data[i] where i is 0..4.package day12
		if you want initilization onther tha 0.0, you can provide a list as follows:
		var data [5] float64 {0.2, 0.8, 1.3, 2.5, 6.9}

		In Go, these arrays are static and their lenght is part of their tyype making them difficult to use.
		Instead, GO has a very similar structure called a slice that is very flexible, dynamic and easy to use. A slice is declared as
		var data[]float64 - so its length is not specified

		We can use an initializing list so var data[] float64{0.2, 0.8, 1.3, 2.5, 6.9}
		will create a five element slice.
*/

// Let us use our convertToKm program with a slice to show how some of this works.

package main

import "fmt"

func main() {
	// len is the length of the slice, the number of elements. Capacity is the number of elements in the underlying array.
	// len is the number of elements in the slice. It is the number of elements you can access.
	// cap is the number of elements in the underlying array. It is the number of elements you can access.
	// The len of the slice is set to the number of elements in the initializer list. The cap is set to the number of elements in the initializer list.
	// The underlying array is created with enough space to hold all the elements.
	var miles = []int{10, 26, 45} //slice len = 3 cap = 3
	var yards = []int{0, 385, 44}
	for i, _ := range miles { // i locally declared _ disregard value

		fmt.Println("Convert Miles and Yards to Kilometers: ", miles[i], "mi. : ", yards[i], "yd.")
		fmt.Printf("Answer is %f kilometers.\n\n", convertToKm(miles[i], yards[i]))

	}
}

func convertToKm(miles int, yards int) float64 {
	kiloMeter := (1.609 * (float64(miles) + float64(yards)/1760.0))
	return kiloMeter
}

/*
	We declared two slices of three elements each and use a for-range statement to loop over it.
	The for range avoids indexing and off by one errors.
	On each iteration the local variable i and the undeclared special variable '_' get the index-value pair
	i is 0 and _ is miles[0] which is thrown away
*/

/*
	Other key ideas for slices: make, len, cap, append
		data := make([]int,100)
	* This creates an int slice of length 100 all of whose elements are 0.
	* Make is a build-it function that can be used to allocate different structured types-so more later.
		len(data) would be 100
	* len() is a built-in function that gives length of slices, strings and channels. 
		data = append (data, 7,8,9,10)
	* This uses the build-in function append to add items to the end of the named slice. Here 4 more values are added so now len(data) is 104
		cap(data)
	* This is the size of the data structure, which can be larger than its len() for reasons of efficiency.
*/
