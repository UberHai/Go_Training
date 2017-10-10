package main

import "fmt"

func main() {
	variable := "UberHai"
	switch variable {
	case "Shawn", "UberHai":
		fmt.Println("What's up grand master flash")
	case "Michael", "Nezsaki":
		fmt.Println("Yooo what up bro??")
	case "Ian", "Lt Dan":
		fmt.Println("Long time no see!")
	case "Tony", "T-Train", "T-Loc":
		fmt.Println("Dude, I'm movin to Texas!!")
	default:
		fmt.Println("Maybe one day..")
	}
}
