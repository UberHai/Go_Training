package main

import "fmt"

func main() {
	if true || false {
		fmt.Println("This will print because TRUE or FALSE evals to true")
	}

	if true && false {
		fmt.Println("This will not run because TRUE and FALSE evals to false because both are not true")
	}
}