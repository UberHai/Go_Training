package main

//Importing is per-file basis
import "fmt"

//Scope is Package level
const p string = "Shawn Wesley Stevens"
const r = "Golang is great!"

func main() {
	
	fmt.Println("r - ", r, " \t")
	//Appears you can revalue const if they have not been called before
	//EDIT It is actually creating a new var r in the main() scope
	//anytime r is referenced after this is not the CONST
	r := "Probably dosnt work"
	r += ", but it does"
	r += " because I was wrong"

	//Scope is main()
	const q = 42

	fmt.Println("p - ", p, " \t")
	fmt.Println("r - ", r, " \t")
	fmt.Println("q - ", q, " \n")
	//Cannot reassign value since it has been called
	//q = 69
	//fmt.Println("q - ", q, " \n")
}
