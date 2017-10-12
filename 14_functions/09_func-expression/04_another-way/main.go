package main

import "fmt"

//Function that returns a function and a string
func makeGreeter() func() string {
	//return a func() that returns a string
	return func() string {
		return "Hello World"
	}
}

func main() {
	//Set variable to function that returns a function and string
	greet := makeGreeter()
	//Call function
	fmt.Println(greet())
}

