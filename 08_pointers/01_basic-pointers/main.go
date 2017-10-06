package main

import "fmt"

//CANNOT TAKE THE MEMORY ADDRESS OF AN IOTA
//const a =  iota

func main() {
	a := 1337
	fmt.Println(a)  // 1337
	fmt.Println(&a) // MEMORY ADDRESS

	//CREATE A VAR b AND SET IT
	//TO TYPE "INT POINTER" AND
	//ASSIGN IT THE MEMORY ADDRESS OF a
	var b *int = &a

	//PRINT OUT WHAT B IS
	fmt.Println(b) // SAME MEMORY ADDRESS AS a
	//*b = DEREFERENCING VAR A
	//IN THIS CASE IT IS THE VALUE OF a
	fmt.Println(*b) // 1337

	//UPDATE THE VALUE AT THIS ADDRESS
	*b = 1101110011110001
	//PRINT OUT VAR a WITH NEW VALUE
	fmt.Println(a)
}
