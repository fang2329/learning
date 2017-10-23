package main

import (
	//      "fmt"
	"testing"
)

const capacity = 1024

func array() [capacity]string {
	var d [capacity]string

	for i := 0; i < len(d); i++ {
		d[i] = "hello"
	}

	return d
}

func slice() []string {
	d := make([]string, capacity)

	for i := 0; i < len(d); i++ {
		d[i] = "world"
	}

	return d
}
func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = array()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = slice()
	}
}
