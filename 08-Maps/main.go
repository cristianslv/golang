package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign key values
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharen"] = "sharen@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Define Map and Assign key values
	emails := map[string]string{
		"Bob":    "bob@gmail.com",
		"Sharen": "sharen@gmail.com",
	}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")

	fmt.Println(emails)
}
