package main

import (
	"fmt"
	"strconv"
)

// define person struct
type Person struct {
	firstName string
	lastName  string
	age       int
	gender    string
}

// greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old."
}

func main() {
	// init person struct
	p1 := Person{"Abhinav", "Robinson", 20, "M"}

	fmt.Println(p1)

	// get single field
	fmt.Println(p1.firstName)

	// change value
	p1.age++

	// value reciever
	fmt.Println(p1.greet())
}
