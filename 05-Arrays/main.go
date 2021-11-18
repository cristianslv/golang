package main

import "fmt"

func main() {
	// Arrays
	// var fruitArray [2]string

	// Assign values
	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"

	// fruitArray := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArray)

	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fruitSlice = append(fruitSlice, "Strawberry")

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
