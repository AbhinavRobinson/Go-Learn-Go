package main

import "fmt"

// Functional Programming
func Adder() func(int) int {
	// internal state
	sum := 0
	// return function which has access to internal state
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// init the function
	sum := Adder()
	for i := 0; i < 10; i++ {
		// modify state with returned function
		fmt.Println(sum(i))
	}
}
