package filter

// Filter construct a new slice from those elements
// for which the function is true
func Filter[S ~[]E, E any](s S, f func(E) bool) S {

	res := S{}

	for _, e := range s {
		if f(e) {
			res = append(res, e)
		}
	}
	return res
}
