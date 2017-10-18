package main

import (
	"encoding/json"
	"fmt"
)
//Lowercase fields wont be exported
//If something comes in as "wisdom score"
//It will put the data into Age field
type Person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"`
	//notExported int
}

func main() {
	//Create a blank person
	var p1 Person
	fmt.Println("First: ",p1.First)
	fmt.Println("Last: " ,p1.Last)
	fmt.Println("Age: ",p1.Age)

	//Create bytes to input directly into the blank person
	//Change age to wisdom score as it will be put into Age as
	//shown in the struct
	bs := []byte(`{"First":"James", "Last":"Bond", "wisdom score":20}`)
	//Input into blank person
	json.Unmarshal(bs, &p1)
	fmt.Println("---------------------")
	fmt.Println("First: ",p1.First)
	fmt.Println("Last: " ,p1.Last)
	fmt.Println("Age: ",p1.Age)
	fmt.Printf("%T\n", p1)
	fmt.Println("---------------------")
}