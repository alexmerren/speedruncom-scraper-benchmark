package main_test

import (
	"fmt"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	for b.Loop() {
		fmt.Println("test")
	}
}
