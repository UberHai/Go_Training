/*
	A row of five black square tiles is to have a number of its tiles replaced with coloured oblong tiles chosen from red (length two), green (length three), or blue (length four).

	If red tiles are chosen there are exactly seven ways this can be done.

	If green tiles are chosen there are three ways.

	And if blue tiles are chosen there are two ways.

	Assuming that colours cannot be mixed there are 7 + 3 + 2 = 12 ways of replacing the black tiles in a row measuring five units in length.

	How many different ways can the black tiles in a row measuring fifty units in length be replaced if colours cannot be mixed and at least one coloured tile must be used?
*/

package main

import "fmt"

func main() {
	solutions := 0
	length := 50
	min := 2
	max := 4
	for i := min; i <= max; i++ {
		solutions += tiling(length, i)
	}
	fmt.Println("Row of length ", length, " fills in ", solutions, " ways")
}

func tiling(m, n int) int {
	solution := 0
	if n > m {
		return solution
	}
	for start := 0; start <= m-n; start++ {
		solution++
		solution += tiling(m-start-n, n)
	}
	return solution
}
