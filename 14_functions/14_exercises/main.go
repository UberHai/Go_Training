package main

import "fmt"

func main() {
	/* EXERCISE 01 */
	myNum := 19
	half, even := divedBy2Even(myNum)
	fmt.Println("Half of myNum = ", half)
	fmt.Println("Is myNum even? = ", even)
	/* END EXERCISE 01 */

	/* EXERCISE 02 */
	divdedExpression := func(a int) (z int, y bool) {
		fmt.Println("Taken in integer: ", a)
		z = a / 2
		if a%2 == 0 {
			y = true
		} else {
			y = false
		}
		return z, y
	}
	h, e := divdedExpression(42)
	fmt.Println("Half of myNum = ", h)
	fmt.Println("Is myNum even? = ", e)
	/* END EXERCISE 02 */

	/* EXERCISE 03 */
	greatest := func(a ...int) int {
		//return index of highest value
		var max, ind int
		for i, v := range a {
			if v > max {
				max = v
				ind = i
			}
		}
		return ind
	}
	data := []int{4,9,23,76,336,2,443,958,2754,1337,9001,69,42,360}
	data05a := greatest(1, 2)
	data05b := greatest(1,2,3)
	data05c := []int{1,2,3,4}
	data05d := greatest()
	//find index of highest value
	ind := greatest(data...)
	//print value
	fmt.Println(data[ind])
	fmt.Println("Ex. 5 - ", data05a, " INDEX")
	fmt.Println("Ex. 5 - ", data05b, " INDEX")
	fmt.Println("Ex. 5 - ", greatest(data05c...), " INDEX")
	fmt.Println("Ex. 5 - ", data05d, " INDEX")
	/* END EXERCISE 03 */

	/* EXERCISE 04 */
	if (true && false) || (false && true) || !(false && false){
		fmt.Println("Should show")
	}
	/* END EXERCISE 04 */

	/* EXERCISE 05 */
	//SEE EXERCISE 02 FOR A 2-IN-1 COMBO
	/* END EXERCISE 05 */
}
/* EXERCISE 01 FUNCTION */
func divedBy2Even(a int) (z int, y bool) {
	fmt.Println("Taken in integer: ", a)
	z = a / 2
	if a%2 == 0 {
		y = true
	} else {
		y = false
	}
	return z, y
}
/* END EXERCISE 01 FUNCTION */
