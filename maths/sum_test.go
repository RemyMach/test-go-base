package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.EqualValues(t, Sum(5, 5), 10, "given two number when they are add we should see the sum!")
}
