package main

import (
	"a"
	"c"
)

func ReadFromNetwork(b []byte) {
	for i, n := 0, len(b); i < n; i++ {
		b[i] = byte(i)
	}
}

func main() {
	a.Buffer.P = 0x12345678
	a.Buffer.L = 100
	a.Buffer.C = 100
	b, _ := c.GetBuffer()
	ReadFromNetwork(b)
}
