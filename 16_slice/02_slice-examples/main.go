package main

import "fmt"

func main() {
	//Setup slice of int variable using make()
	mySlice := make([]int, 0, 2)

	fmt.Println(("------------------------"))
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println(("------------------------"))

	for i := 0; i < 64; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), "\tCapacity: ", cap(mySlice), "\tValue: ", mySlice[i])
	}

	//Doubling method for capacity stops around 10,000
	//After that it changes to +20%
	longSlice := make([]int, 0, 100)
	for i :=0; i < 12291; i++ {
		longSlice = append(longSlice, i)
		if (i+1) % 10 == 0 {
			fmt.Println("Len: ", len(longSlice), "\tCapacity: ", cap(longSlice), "\tValue: ", longSlice[i])
		}
	}
	//How to joyfully append slices together
	yourSlice := []string{"Monday", "Tuesday"}
	yourOtherSlice := []string{"Wednesday", "Thursday", "Friday"}
	yourSlice = append(yourSlice, yourOtherSlice...)
	fmt.Println(yourSlice)
	/////////////////////////////////////////////////////
	//How to joyfully delete an element from a slice
	//"Get me items up to but before 2 [0,1] and items 3 and beyond [3,...]
	yourSlice = append(yourSlice[:2], yourSlice[3:]...)
	fmt.Println(yourSlice)
}
