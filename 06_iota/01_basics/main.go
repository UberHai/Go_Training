package main

import "fmt"

const (
	// Will equate to a = 0, b = 1, c = 2
	a = iota
	b = iota
	c = iota
)

const (
	// iota resets on new scope declaration
	// Will equate to d = 0, e = 1, f = 2
	d = iota
	e
	f
)

const (
	_  = iota             // 0 - throw away variable
	kb = 1 << (iota * 10) // 1 << (1 * 10)
	mb = 1 << (iota * 10) // 1 << (2 * 10)
	gb = 1 << (iota * 10) // 1 << (3 * 10)
	tb = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	fmt.Println("IOTA TESTING")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println("kb =\t", kb)
	fmt.Println("mb =\t", mb)
	fmt.Println("gb =\t", gb)
	fmt.Println("tb =\t", tb)
}
