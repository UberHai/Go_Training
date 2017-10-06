package main

import "fmt"

func main() {
	x := 1337 % 69 //66 Remainder 17
	fmt.Println("Remainder of 69 divided into 1337: ", x)

	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println(i, " : ODD")
		} else {
			fmt.Println(i, " : EVEN")
		}
	}

	//IS THE NUMBER... PRIME?
	n := -1

	for n != 0 {
		flag := 0
		fmt.Print("Enter a positive integer: ")
		fmt.Scanf("%d", &n)

		for i := 2; i <= n/2; i++ {
			// condition for nonprime number
			if n%i == 0 {
				fmt.Println(n)
				fmt.Println(n, "%", i)
				flag = 1
				break
			}
		}

		if flag == 0 {
			fmt.Printf("%d is a prime number.\n", n)
		} else {
			fmt.Printf("%d is not a prime number.\n", n)
		}
	}
}
