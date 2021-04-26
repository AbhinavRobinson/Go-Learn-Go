package main

import "fmt"

func main() {
	ids := []int{23, 42, 34, 12, 54, 73}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
}
