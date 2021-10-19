package main

import "fmt"

type student interface {
	stdinfo()
}

type stud struct {
	name string
	age  int
}

func (s stud) stdinfo() {
	fmt.Printf("%T format,%v value", s, s)
}
func main() {
	var n student
	std1 := stud{"mahesh", 12}
	n = std1
	n.stdinfo()

}
// package main

// import "fmt"

// //empty interface
// func want(i interface{}) {
// 	fmt.Printf("Type = %T,value = %v\n", i, i)

// }

// func assert(i interface{}) {
// 	key, value := i.(int)
// 	fmt.Println(key, value)

// }
// func main() {
// 	s := "mahesh siva"
// 	want(s)
// 	a := 89
// 	want(a)

// 	std1 := struct {
// 		name string
// 	}{
// 		name: "mahesh",
// 	}
// 	want(std1)

// 	//type assertion
// 	var n interface{} = 34
// 	assert(n)
// 	var k interface{} = "string"
// 	assert(k)

// }
