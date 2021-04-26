package main

import "fmt"

// define person struct
type Person struct {
	firstName string
	lastName  string
	age       int
	gender    string
}

func main() {
	// init person struct
	p1 := Person{"Abhinav", "Robinson", 20, "M"}

	fmt.Println(p1)

	// get single field
	fmt.Println(p1.firstName)

	p1.age++
	fmt.Println(p1)
}
