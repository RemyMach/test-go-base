package benchmark

import (
	"testing"
)

func Sum(x int, y int) int {
	return x + y
}


func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(5, 7)
	}
}
