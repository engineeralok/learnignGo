package main

import "fmt"

func main() {
	var lambai, chaudhai, kshetrafal float32
	fmt.Print("Lambai ki value likhiye : ")
	fmt.Scanf("%f\n", &lambai)
	fmt.Print("Chaudhai ki value likhiye : ")
	fmt.Scanf("%f", &chaudhai)
	kshetrafal = lambai * chaudhai
	fmt.Print("Ayat ka kshetrafal : ", kshetrafal, "meter square. Jiski lambai hai : ", lambai, " meter. Aur chaudhai hai : ", chaudhai, " meter")
}
