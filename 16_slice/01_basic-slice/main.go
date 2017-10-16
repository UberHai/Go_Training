package _1_basic_slice

import "fmt"

func main() {
	//Create a slice of ints
	firstSlice := []int{1, 1, 2, 3, 5, 8, 13, 21}
	//What is its length? 8
	fmt.Println(len(firstSlice))
	//What is it's capacity? 8
	fmt.Println(cap(firstSlice))
	//What is it's type? []int
	fmt.Printf("%T\n", firstSlice)
	//Show the slice
	fmt.Println(firstSlice)
	//append a value to position 9.. longer than its capacity
	firstSlice = append(firstSlice, 34)
	//New length since append? 9
	fmt.Println(len(firstSlice))
	//New capacity since append? 16 ( it doubles )
	fmt.Println(cap(firstSlice))
	//Show new array
	fmt.Println(firstSlice)
}
