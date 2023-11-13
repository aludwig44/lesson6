package main

import (
	"fmt"
	"math/rand"
)

func main() {
	piggybank := 0.0
	count := 0
	for piggybank <= 2000 {
		coin := rand.Intn(3)
		switch coin {
		case 0:
			piggybank += 5
			count += 1
			fmt.Printf("You've deposited a nickel! Your piggybank total is $%.2f.\n", (piggybank / 100))
		case 1:
			piggybank += 10
			count += 1
			fmt.Printf("You've deposited a dime! Your piggybank total is $%.2f.\n", (piggybank / 100))
		case 2:
			piggybank += 25
			count += 1
			fmt.Printf("You've deposited a quarter! Your piggybank total is $%.2f.\n", (piggybank / 100))
		}
	}
	fmt.Printf("Your piggypank total is $%.2f and contains %v coins.  ", (piggybank / 100), count)
}
