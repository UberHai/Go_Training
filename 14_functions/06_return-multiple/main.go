package main

import "fmt"

func main() {
	fmt.Println(greet("John", "Doe"))
}

//Return multiple values by putting the types inside parens after closing parameters
func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname), fmt.Sprint(lname)
}
