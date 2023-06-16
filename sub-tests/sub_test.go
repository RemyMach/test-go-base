package sub_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Sum(x int, y int) int {
	return x + y
}

func TestSum(t *testing.T) {
	cases := []struct {
		x, y, want int
	}{
		{5, 5, 10},
		{0, 0, 0},
		{2, 2, 4},
		{-5, -5, -10},
	}

	for _, c := range cases {
		got := Sum(c.x, c.y)
		assert.Equal(t, c.want, got, "Sum(%d, %d) should be %d, but got %d", c.x, c.y, c.want, got)
	}
}
