package main

import "fmt"

func bill(val *int) {
	*val = 89
}

func void() *int {
	c := 56
	return &c

}
func main() {
	//creating pointers
	a := 3
	b := &a
	fmt.Println(b)
	//dereferencing thn pointer
	fmt.Println(*b)
	//passing pointer to a function
	bill(b)
	fmt.Println(a)

	c := new(int)
	fmt.Printf("%T \n", c)
	*c = 80
	fmt.Println(*c)
	d := void()
	fmt.Print(*d)

}
