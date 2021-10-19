package main

import (
	"fmt"
)

func main() {
	name := "Naveen"
	runes := []rune(name)

	for i := 0; i < len(runes); i++ {
		defer fmt.Printf("%c ", runes[i])
	}
	// runes := []rune(s)
	// for i := 0; i < len(runes); i++ {
	//     fmt.Printf("%c ", runes[i])
	// }
}
