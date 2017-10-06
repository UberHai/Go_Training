package main

import "fmt"

func zero(x int) {
	x = 0
	fmt.Printf("Memory Address of the x in function : %p\n", &x)
}

func zeroWorking(y *int) {
	*y = 0
	fmt.Printf("Memory Address of the y in function : %p\n", &y)
}

func main() {
	x := 5
	zero(x) //Set X to zero through a function
	fmt.Printf("Memory Address of the x passed in : %p\n", &x)
	fmt.Println(x) //Just kidding, it dosn't work like that. X is still 5
	z := 100
	zeroWorking(&z) // Hell this just might work
	fmt.Printf("Memory Address of the z passed in : %p\n", &z)
	fmt.Println(z) // Annnnnd it does
	fmt.Printf("Memory Address of the z changed to 0 : %p\n", &z)

}
