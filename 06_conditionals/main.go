package main

import "fmt"

func main() {
	x := 15
	y := 10

	// if else if else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x == y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is equal to %d\n", y, x)
	}

	color := "red"
	// switch
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is red")
	default:
		fmt.Println("Color is rnot blue or red")
	}
}
