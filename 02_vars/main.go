package main

import "fmt"

func main() {
	name, email := "Abhinav", "abhinav@designs.studio"
	var age int8 = 20
	const isDev = true
	size := 1.3

	fmt.Println(name, email, age, isDev, size)
	fmt.Printf("%T ", name)
	fmt.Printf("%T ", email)
	fmt.Printf("%T ", age)
	fmt.Printf("%T ", isDev)
	fmt.Printf("%T ", size)
}
