package main

import "fmt"

func main() {
	a := 1337
	b := "Elite"

	fmt.Println("a -\t", a)
	fmt.Println("b -\t", b)
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("%d\n", &a)
	fmt.Println("b's memory address - ", &b)
	fmt.Printf("%d\n", &b)
}
