package main

import (
	"fmt"
	"strconv"
)

// define person struct
type Person struct {
	// firstName string
	// lastName  string
	// age       int
	// gender    string
	firstName, lastName, gender string
	age                         int
}

// greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old."
}

// has Bday (ptr reciever)
func (p *Person) hasBday() {
	p.age++
}

// getMarried (ptr reciever)
func (p *Person) getMarried(spouceLastName string) {
	if p.gender == "F" {
		p.lastName = spouceLastName
	} else {
		return
	}
}

func main() {
	// init person struct
	p1 := Person{"Abhinav", "Robinson", 20, "M"}

	fmt.Println(p1)

	// get single field
	fmt.Println(p1.firstName)

	// change value with ptr reciever
	p1.hasBday()

	// create another person
	p2 := Person{"Samantha", "Smith", 20, "F"}
	fmt.Println(p2.greet())

	// value reciever
	fmt.Println(p1.greet())

	p1.getMarried("Smith")
	p2.getMarried("Robinson")

	fmt.Println(p2.greet())
}
