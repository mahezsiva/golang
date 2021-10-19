package main

import (
	"errors"
	"fmt"
)

func bell(a int) (int, error) {
	if a < 0 {
		return 0, errors.New("value is negative")
	}
	return a, nil
}

func main() {

	a := -10
	c, err := bell(a)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(c)

}
