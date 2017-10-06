package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	//Print memory address before input
	fmt.Printf("Meters Memory Address : %d\n", &meters)
	fmt.Print("Enter Meters Swam : ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	//Print memory address after input
	fmt.Printf("Meters Memory Address : %d\n", &meters)
	fmt.Println(meters, " meters is ", yards, " yards.")
}
