package main

import "fmt"

func main() {
	//Create functions in other functions with function expressions
	//greeting is a variable that is an anonymous function
	greeting := func() {
		fmt.Println("Hello World")
	}
	//Call greeting as a function : greeting()
	greeting()
	fmt.Printf("%T\n", greeting)
}

