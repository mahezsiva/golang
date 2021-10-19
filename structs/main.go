package main

import "fmt"

func main() {

	std3 := struct {
		firstname string
		age       int
	}{
		firstname: "diya",
		age:       12,
	}
	fmt.Println(std3)
}
