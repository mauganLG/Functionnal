package reduce

// Reduce apply a function of two arguments to the elements of a slice
// from left to right to reduce the slice to a one value
func Reduce[S ~[]E, E any](s S, f func(current, next E) E) E {

	var res E
	for _, e := range s {
		res = f(res, e)
	}

	return res
}
