package main

import "fmt"

func main() {
	fmt.Println(greet("John", "Doe"))
}

//You can name returns by putting the name and type inside parenthesis after closing the parameters
func greet(fname, lname string) (s string){
	s = fmt.Sprint(fname, lname)
	return
}