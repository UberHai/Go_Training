package main

import "fmt"

func main() {
	//Character A represents an alias of int32
	//The letter A value is 65 according to ASCII/UTF-8
	letter := 'A'
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)
	//You can grab characters/runes from inside strings by using brackets.
	//The position 3 of Apple is l.
	//The letter l has a value of 108 according to ASCII/UTF-8
	//Position 3 of a string is a character/rune, which is an alias of int32
	letter = rune("Apple"[3])
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)
	//Start a for loop from the character A until character z
	//Display i, it's character reference as a string, and the bucket it will placed in.
	for i := 65; i <= 122; i++ {
		fmt.Println(i, " - ", string(i), " - ", i % 12)
	}

}
