package c

import "a"

func F() uint32 {
	x := a.A
	y := uint32(12)
	for y != 0 {
		t := y
		y = x % y
		x = t
	}
	return x
}
