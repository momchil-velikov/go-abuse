package c

import (
	"a"
	"math/rand"
)

func GetBuffer() ([]byte, int) {
	return a.Buffer.B, rand.Int()
}
