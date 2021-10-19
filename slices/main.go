package main

import "fmt"

func main() {

	//creating slice
	a := []int{1, 2, 3, 4, 5}
	b := a[1:4]

	fmt.Println(b)
	//modifying th slice
	for i := range b {
		b[i]--
	}
	fmt.Println(a)
	fmt.Println(b)

	//copying entire content of slice
	c := a[:]
	fmt.Println(c)

	//length and capacity of slice
	fmt.Println("length=", len(c), "capacity=", cap(c))

	//cretaing slice  using make
	i := make([]int, 3, 4)
	fmt.Println(i)

	//appending slices
	s := []string{"ajay", "akash", "aswin"}
	fmt.Println("before appending=", s)
	s = append(s, "siva")
	fmt.Println("after appending=", s)

	//appending two slices to one

	f := []string{"dinesh", "vikram"}
	g := append(s, f...)
	fmt.Println("both slices=", g)

}
