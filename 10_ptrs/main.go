package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T %T\n", a, b)

	// read value in *
	fmt.Println(a, *b)

	// change value with pointer
	*b = 10
	fmt.Println(a, *b)
}
