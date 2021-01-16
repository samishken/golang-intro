package main

import "fmt"

func main() {
	// This is an example of a string
	fmt.Println("Hello")

	var ourString string
	ourString = "This is a Test"
	fmt.Println(ourString)

	//This is an example of a Rune
	var ourRune rune
	ourRune = 'g'
	fmt.Printf("Type: %T, \v Value: %v\n", ourRune, ourRune)

	// Example of needing special character
	fmt.Println("Don't touch That, that's my food")
	fmt.Println("She said, \"Thats my food too\"")

}
