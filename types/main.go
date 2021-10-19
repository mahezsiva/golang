package main

import "fmt"

func main() {

	//complex types
	a := 2 + 3i
	b := 9 + 5i
	fmt.Println(a + b)
	fmt.Println(a - b)
	//string type
	first := "mahesh"
	last := "siva"
	fmt.Println(first + " " + last)
	//type conversion
	var A int = 10
	B := 10.5
	C := A + int(B)
	fmt.Print(C)
	//signed Intergers
	fmt.Printf("type of A is %T\ntype of B is%T\ntype of C is %T\n", A, B, C)

}
