package main

import "fmt"

func main() {
	// // arrays
	// var fruitArr []string

	// // assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// declare and assign
	fruitArr := []string{"Apple", "Orange"}
	// print array
	fmt.Println(fruitArr)

	// append
	fruitArr = append(fruitArr, "Grape")
	fmt.Println(fruitArr)
}
