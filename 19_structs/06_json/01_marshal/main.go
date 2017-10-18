package main

import (
	"encoding/json"
	"fmt"
)
//Lowercase fields wont be exported
type Person struct {
	First string
	Last  string
	Age   int
	notExported int
}

func main() {
	//Create a person
	p1 := Person{"James", "Bond", 20, 007}
	//Marshal the person data
	bs, _ := json.Marshal(p1)
	//Print the marshaled data
	fmt.Println(bs)
	//Print type
	fmt.Printf("%T \n", bs)
	//Print the data as a string
	fmt.Println(string(bs))
}