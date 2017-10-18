package main

import "fmt"

//Create a new type
//It has an underlying type of int
type foo int

func main() {
	//Create a new var foo
	var myAge foo
	myAge = 55
	fmt.Printf("%T %v \n", myAge, myAge)
	//Create a new var int
	var yourAge int
	yourAge = 29
	fmt.Printf("%T %v \n", yourAge, yourAge)


	/* This wont work because you add multiple types
	fmt.Println(myAge + yourAge)
	*/
	//You must convert the type before attempting
	fmt.Println(int(myAge) + yourAge)
	}