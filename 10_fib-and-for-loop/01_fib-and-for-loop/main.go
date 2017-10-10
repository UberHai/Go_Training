package _1_fib_and_for_loop

import "fmt"

// Fib returns the nth number in the Fibonacci series.
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func main() {
	//Print the 8th number of Fibonacci series
	fmt.Println(Fib(8))

	//Print outer and inner loop sequence
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}
}
