package main

import "fmt"

func main() {
    toppings := [3]string {"cheese", "onions", "olives"}

	fmt.Println("Our choice of toppings include:")

	// i = index
    for i, value := range toppings {
        fmt.Println(i + 1, value)
    }
}