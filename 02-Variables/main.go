package main

import "fmt"

func main() {
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	var age int = 19
	var isCool = true
	isCool = false

	// Shorthand
	// name := "Cristian Eduardo"
	// email := "cristian@gmail"

	name, email := "Cristian", "cris@outlook"

	size := 1.3

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", isCool)
}
