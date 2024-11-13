package filter

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterInt(t *testing.T) {

	sliceInt := []int{0, 1, 2, 3, 4, 5}

	trueEvenSliceInt := []int{0, 2, 4}
	even := func(i int) bool {
		return i%2 == 0
	}

	evenSliceInt := Filter(sliceInt, even)

	res := reflect.DeepEqual(evenSliceInt, trueEvenSliceInt)

	assert.EqualValues(t, true, res)
}

func TestFilterFloat(t *testing.T) {

	sliceFloat := []float32{0, 1.1, 2.2, 3.3, 4.4, 5.5}

	trueSliceFloat := []float32{1.1, 2.2}
	filt := func(i float32) bool {

		if i > 1.0 && i < 3.0 {
			return true
		}
		return false
	}

	fSliceInt := Filter(sliceFloat, filt)

	res := reflect.DeepEqual(fSliceInt, trueSliceFloat)

	assert.EqualValues(t, true, res)
}

func TestFilterFloatDecimal(t *testing.T) {

	sliceFloat := []float64{0.199, 1.111, 2.265, 3.343, 4.478, 5.132}

	trueSliceFloat := []float64{0.199, 1.111, 5.132}
	filtDecimalOne := func(i float64) bool {

		d := int(i*10) % 10

		if d == 1 {
			return true
		}
		return false
	}

	fSliceInt := Filter(sliceFloat, filtDecimalOne)

	res := reflect.DeepEqual(fSliceInt, trueSliceFloat)

	assert.EqualValues(t, true, res)
}

func TestFilterRune(t *testing.T) {

	sliceFloat := []rune{'a', 'A', '@', 'b', 'B', '&'}

	trueSliceFloat := []rune{'a', 'b'}
	filtDecimalOne := func(r rune) bool {

		if r >= 'a' && r <= 'z' {
			return true
		}

		return false
	}

	fSliceInt := Filter(sliceFloat, filtDecimalOne)

	res := reflect.DeepEqual(fSliceInt, trueSliceFloat)

	assert.EqualValues(t, true, res)
}
