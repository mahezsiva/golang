package main

import "fmt"

// creating variadic functions
func vard(a int, b ...int) {

	fmt.Print(a)
	fmt.Print(b)
}

func main() {
	vard(1, 2, 3, 4, 5, 6)

}
