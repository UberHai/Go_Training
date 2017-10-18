package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

//func [receiver] name([parameter...]) return
//This function is received by the person struct. Returns a string
func (p person) fullName() string {
	return p.first + " " + p.last
}

func main() {
	var p3 person
	p3.first = "Elon"
	p3.last = "Musk"

	fmt.Println(p3.fullName())
}
