package main

import "fmt"

func main() {
	//String array of length 58
	var x [58]string

	//for loop through ascii #65 thru #122
	//assign i less 65 of x to string(i)
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	//Print array
	fmt.Println(x)
	//Print array[42]
	fmt.Println(x[42])

	//Create var y as int array 256 long
	var y [256]int
	//Print the length of y
	fmt.Println(len(y))
	//Print y[42]
	fmt.Println(y[42])
	for i:= 0; i < len(y);i++ {
		y[i] = i
	}
	for _, v := range y {
		fmt.Printf("%v - %T -\t %b\n", v, v, v)
	}
}
