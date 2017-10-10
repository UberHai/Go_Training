package main

import "fmt"

func main() {
	if true {
		fmt.Println("This ran because it was true.")
	} else {
		fmt.Println("This did not run because the first condition is true.")
	}

	if false {
		fmt.Println("Totally will not run because it is false.")
	}

	// ! Exclamation Point
	if !true {
		fmt.Println("This did not run because it was not true.")
	} else if !false {
		fmt.Println("This did run because the other statement was not true and this one was not false.")
	} else {
		fmt.Println("What is neither true nor false?")
	}

	if !false {
		fmt.Println("Totally will run because it is not false.")
	}

	//Testing
	for i :=0; i < 14;i++ {
		if i / 3 == 4 {
			//Stop after reaching 12
			fmt.Println("Twelve is the magic number.")
			break
		} else if i % 3 == 0 {
			//Print when i is divisible wholly by 3
			fmt.Println("Divisible by three")
		} else if i == 13 {
			//Wont print.
			fmt.Println("Unlucky")
		}
	}
}
