package Map

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapInt(t *testing.T) {

	sliceInt := []int{0, 1, 2, 3, 4, 5}

	trueDoubleSliceInt := []int{0, 2, 4, 6, 8, 10}
	double := func(i int) int {
		return i * 2
	}

	doubleSliceInt := Map(sliceInt, double)

	res := reflect.DeepEqual(doubleSliceInt, trueDoubleSliceInt)

	assert.EqualValues(t, true, res)
}

func TestMapIntFalse(t *testing.T) {

	sliceInt := []int{0, 1, 2, 3, 4, 5}

	trueDoubleSliceInt := []int{0, 2, 4, 6, 8}
	double := func(i int) int {
		return i * 2
	}

	doubleSliceInt := Map(sliceInt, double)

	res := reflect.DeepEqual(doubleSliceInt, trueDoubleSliceInt)

	assert.EqualValues(t, false, res)
}

func TestMapString(t *testing.T) {

	sliceStr := []string{"a", "b", "c", "d"}

	trueSliceStr := []string{"A", "B", "C", "D"}

	upperSliceInt := Map(sliceStr, strings.ToUpper)

	res := reflect.DeepEqual(upperSliceInt, trueSliceStr)

	assert.EqualValues(t, true, res)
}

func TestMapStringDouble(t *testing.T) {

	sliceStr := []string{"a", "b", "c", "d"}

	trueSliceStr := []string{"aa", "bb", "cc", "dd"}

	doubleStr := func(s string) string {
		return s + s
	}

	upperSliceInt := Map(sliceStr, doubleStr)

	res := reflect.DeepEqual(upperSliceInt, trueSliceStr)

	assert.EqualValues(t, true, res)
}
