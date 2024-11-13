package reduce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduceInt(t *testing.T) {

	sliceInt := []int{0, 1, 2, 3, 4, 5}

	trueRec := 15
	sum := func(i, j int) int {
		return i + j
	}

	res := Reduce(sliceInt, sum)

	assert.EqualValues(t, trueRec, res)
}

func TestReduceString(t *testing.T) {

	sliceStr := []string{"a", "b", "c", "d"}

	trueStr := "abcd"

	concat := func(i, j string) string {
		return i + j
	}
	finalStr := Reduce(sliceStr, concat)

	assert.EqualValues(t, trueStr, finalStr)
}
