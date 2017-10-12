package main

import "fmt"

func main() {
	//Create array of float64
	data := []float64{43, 57, 87, 25, 54, 67}
	//Call average with a variadic argument. Data could be 1, 2, or 100+ numbers long
	n := average(data...)
	fmt.Println(n)
}

//Create a func average that takes any number of parameters (variadic) that are float64
//Return a float64
func average(sf ...float64) float64 {
	//Print the sf variable
	fmt.Println(sf)
	//Print the type of sf []float64
	fmt.Printf("%T \n", sf)
	//Create an accumulator
	var total float64
	//For range loop through each item in sf
	//_ blank identifier used to throw away id from range
	//v for the value returned from range
	for _, v := range sf {
		total += v
	}
	//Print how many arguments are taken in
	fmt.Println(len(sf))
	return total / float64(len(sf))
}
