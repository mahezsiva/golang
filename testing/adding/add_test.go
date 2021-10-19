package add_test

import (
"testing"
""
)

type plus struct{
	x int
	y int
	expected int
}

func testact(t *testing.T){

	plus:=plus{
		x: 2,
		y: 2,
	}

	plus.expected=add.act(plus.x,plus.y)

}