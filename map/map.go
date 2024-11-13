package Map

// Map apply a function to every element of the slice
// and return a new slice of it
func Map[S ~[]E, E any](s S, f func(E) E) S {
	res := make(S, len(s))

	for e := range s {
		res[e] = f(s[e])
	}
	return res
}
