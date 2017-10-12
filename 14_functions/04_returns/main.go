package main

import "fmt"

func main() {
	//Call greet with two arguments
	fmt.Println(greet("Jane", "Doe"))
}

//Declare greet with 2 parameters
//return fname + " " + lname as a string
func greet(fname, lname string) string {
	return fmt.Sprint(fname, " ", lname)
}
