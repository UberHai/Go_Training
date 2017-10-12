package main

import "fmt"

func main() {
	fmt.Println(factorial(10))
}

//function that takes an int and returns an int
func factorial(x int) int {
	//return 1 if x = 0
	if x == 0 {
		return 1
	}
	//return the value x times the return from function(x-1)
	//keeps running until x = 0
	return x * factorial(x-1)
}
