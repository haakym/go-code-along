package main

import "fmt"

func main() {
    numbers := []int {1, 2, 3, 4, 5}

	numSlice1 := numbers[3:5]
	numSlice2 := numbers[:2]
	numSlice3 := numbers[2:]

	fmt.Println("numSlice1 = ", numSlice1)
	fmt.Println("numSlice2 = ", numSlice2)
	fmt.Println("numSlice3 = ", numSlice3)

	numSlice4 := make([]int, 5, 10)

	fmt.Println("numSlice4 = ", numSlice4)

}