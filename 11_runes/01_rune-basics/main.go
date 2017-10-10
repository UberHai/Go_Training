package main

import "fmt"

func main() {
	fmt.Println([]byte("Hello"))

	for i := 5150; i <= 5250; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

	//Single quotes indicating it is a single rune
	foo := 'z'
	fmt.Println(foo)
	//Print the type of 'z'
	fmt.Printf("%T \n", foo)
	//Rune is an alias for Int32, so the type wont change
	fmt.Printf("%T \n", rune(foo))

}
