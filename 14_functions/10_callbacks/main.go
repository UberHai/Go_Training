package main

import "fmt"

func visit(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}
func main() {
	//The func(n int) is called back by the visit() above
	xs := visit([]int{1, 1, 2, 3, 4, 1, 5}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs) //2, 3, 4, 5
}
