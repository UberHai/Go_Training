package main

import "fmt"

func main() {
	variable := "Shawn"
	switch variable {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Jordan":
		fmt.Println("Whazzup Jordan")
	case "McKlusky":
		fmt.Println("What's Up McKlusky")
	case "Shawn":
		fmt.Println("Wazzup Shawn")
	default:
		fmt.Println("I don't quite know who you are.")
	}
}
