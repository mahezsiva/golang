package main

import "fmt"

func main() {

	//untyped constants
	const a = "mahesh"
	var h = a
	fmt.Println(h)
	//typed constants
	const b = "siva"
	type myString string
	const value myString = b
	fmt.Print(value)
	//numeric constants
	const c = 5
	var Var int = c
	fmt.Printf("%T", Var)
}
