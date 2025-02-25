package random_number

import (
	"fmt"
	"math/rand"
	"time"
)

func Run() {
	// Input entries
	var range_start, range_end int
	fmt.Println("Enter the range start:")
	fmt.Scanln(&range_start)
	fmt.Println("Enter the range end:")
	fmt.Scanln(&range_end)

	// Define seed
	rand.Seed(time.Now().UnixNano())

	// Generate random number
	number_random := rand.Intn(range_end-range_start+1) + range_start

	// Loop until the number is found

	fmt.Println("Guest random number")
	var number_guest int

	for number_random != number_guest {
		if number_random > number_guest {
			fmt.Println("The number is higher")
		} else {
			fmt.Println("The number is lower")
		}
		fmt.Scanln(&number_guest)
	}

	fmt.Println("Congratulations! You found the number")

}
