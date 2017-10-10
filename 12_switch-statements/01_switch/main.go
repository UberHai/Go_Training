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
		//fallthrough to the next
		fallthrough
	case "Wesley":
		fmt.Println("Wesley")
		//fallthrough to the next
		fallthrough
	case "Stevens":
		fmt.Println("Stevens")
		//end here -- no fallthrough
	default:
		fmt.Println("I don't quite know who you are.")
	}
}
