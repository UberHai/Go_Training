package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)
//Lowercase fields wont be exported
//If something comes in as "wisdom score"
//It will put the data into Age field
type Person struct {
	First 		string
	Last  		string
	Age   		int `json:"wisdom score"`
	notExported int
}

func main() {
	//Create a blank person
	p1 := Person{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1)

	fmt.Println("---------------------")
	fmt.Println("First: ",p1.First)
	fmt.Println("Last: " ,p1.Last)
	fmt.Println("Age: ",p1.Age)
	fmt.Printf("%T\n", p1)
	fmt.Println("---------------------")

	//Creating an empty Person
	var p2 Person
	//Create a reader holding the values to put into p2
	rdr := strings.NewReader(`{"First":"Shawn", "Last":"Stevens", "wisdom score":26}`)
	//Decode the values in the reader and insert into the
	//Pointed memory of p2
	json.NewDecoder(rdr).Decode(&p2)
	//Print to test if code worked correctly
	fmt.Println("---------------------")
	fmt.Println("First: ",p2.First)
	fmt.Println("Last: " ,p2.Last)
	fmt.Println("Age: ",p2.Age)
	fmt.Printf("%T\n", p2)
	fmt.Println("---------------------")
}