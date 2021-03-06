package main

import "fmt"

func main() {
	ids := []int{23, 42, 34, 12, 54, 73}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// W/o index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum of ids is %d\n", sum)

	// Range with map
	emails := map[string]string{"Abhinav": "abhinav@designs.studio", "Joe": "Joe@JoeDonuts.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
}
