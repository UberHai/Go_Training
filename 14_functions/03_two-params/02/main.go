package main

import "fmt"

func main() {
	greet("John", "Doe")
}

//Same type variables can share the same type declaration. Put it after all parameters are named.
func greet(fname, lname string) {
	fmt.Println(fname, lname)
}
