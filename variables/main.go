package main

import (
	"fmt"
	"math"
)

func main() {
	//declaring a variable
	var a int = 9
	fmt.Println(a)
	//declaring multiple variables
	var c, b int = 2, 3
	fmt.Println(c, b)
	//short hand declaration
	//typr inference
	d := 9
	e := 10
	fmt.Println(d + e)
	//declaring multiple values using short hand
	A, B := 10.8, 59.8
	C := A + B
	fmt.Println(C)
	D := math.Min(A, B)
	fmt.Println(D)
	E := math.Max(A, B)
	fmt.Println(E)

}
