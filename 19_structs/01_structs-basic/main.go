package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"James", "Bond", 42}
	p2 := person{"Mary", "Alice", 38}
	var p3 person
	p3.first	 = "Elon"
	p3.last		 = "Musk"
	//p3.age set to zero value, for int is 0
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println(p3.first, p3.last, p3.age)
}
