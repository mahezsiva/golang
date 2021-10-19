package main

import "fmt"

func main() {
	//multiple switch statements
	switch letter := "i"; letter {

	case "i", "a", "e", "u":
		fmt.Println("vowels")
	default:
		fmt.Println("constonents")

	}
}
