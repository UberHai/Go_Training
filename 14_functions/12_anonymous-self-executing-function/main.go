package main

import "fmt"

func main() {
	//Create a function anonymously.. and have it call itself
	func() {
		fmt.Println("I'm Driving!")
	}()
}