package main

import "fmt"

func main() {

	age := 18

	if age >= 16 {
		fmt.Println("You Can Drive")
	} else {
		fmt.Println("You Can't Drive")
	}

	// You can use else if to perform different actions,
	// but once a match is reached the rest of the conditions aren't checked

	if age >= 16 {
		fmt.Println("You Can Drive")
	} else if age >= 18 {
		fmt.Println("You Can Vote")
	} else {
		fmt.Println("You Can Have Fun")
	}

	// Switch statements are used when you have limited options

	switch age {
		case 16: fmt.Println("Go Drive")
		case 18: fmt.Println("Go Vote")
		default: fmt.Println("Go Have Fun")
	}

}