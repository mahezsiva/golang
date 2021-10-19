package main

import "fmt"

func main() {

	//creating an array
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)
	//short hand declartion
	A := [...]int{1, 2, 3, 4, 5}
	fmt.Println(A, len(A))
	//Iterating using range
	D := []int{23, 43, 2, 4, 5}
	for i, j := range D {
		fmt.Printf("index %d and its value %v\n", i, j)
	}
	//blank identifier
	for _, k := range D {
		fmt.Printf("value %v\n", k)
	}

}
