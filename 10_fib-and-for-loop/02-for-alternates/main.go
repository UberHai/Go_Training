package main

import "fmt"

func main() {

	//Condition in declaration but init and post outside
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}
	///////////////////////////////////////////////////////////////

	k := 0
	for {
		fmt.Println(k)
		if k >= 10 {
			//Will stop this loop when K is bigger than 10
			break
		}
		k++
	}
	///////////////////////////////////////////////////////////////

	l := 0
	for {
		//Increment L at beginning
		l++
		//If the remainder of l divided by 2 is 0 continue to next statement
		if l%2 == 0 {
			continue
		}
		//Printing l is the next statement
		fmt.Println(l)
		//Condition so it dosn't run forever.. if higher than 50 break
		if l >= 50 {
			break
		}
	}
	/////////////////////////////////////////////////////////////////
	//No condition for loop. Will run forever. Ctrl+C to stop
	//i := 0
	//for {
	//	fmt.Println(i)
	//	i++
	//}
}
